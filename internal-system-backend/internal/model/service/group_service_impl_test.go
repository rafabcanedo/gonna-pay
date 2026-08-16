package service_test

import (
	"errors"
	"testing"

	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/mocks"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/service"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGroupService_Create(t *testing.T) {
	t.Run("success without members", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		created := testutil.NewGroupFixture()
		group := domains.NewGroup("user-1", "Viagem", "Travel")
		mockRepo.EXPECT().Create(group, []string{}).Return(created, nil)

		result, err := svc.Create(group, []string{})

		assert.NoError(t, err)
		assert.Equal(t, "group-1", result.ID)
	})

	t.Run("success with members", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		created := testutil.NewGroupFixture()
		group := domains.NewGroup("user-1", "Viagem", "Travel")
		memberIDs := []string{"contact-1", "contact-2"}

		mockRepo.EXPECT().IsContactOwnedBy("contact-1", "user-1").Return(true, nil)
		mockRepo.EXPECT().IsContactOwnedBy("contact-2", "user-1").Return(true, nil)
		mockRepo.EXPECT().Create(group, memberIDs).Return(created, nil)

		result, err := svc.Create(group, memberIDs)

		assert.NoError(t, err)
		assert.Equal(t, "group-1", result.ID)
	})

	t.Run("forbidden - contact not owned by user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		group := domains.NewGroup("user-1", "Viagem", "Travel")
		mockRepo.EXPECT().IsContactOwnedBy("contact-1", "user-1").Return(false, nil)

		_, err := svc.Create(group, []string{"contact-1"})

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("repo error on IsContactOwnedBy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		group := domains.NewGroup("user-1", "Viagem", "Travel")
		mockRepo.EXPECT().IsContactOwnedBy("contact-1", "user-1").Return(false, errors.New("db error"))

		_, err := svc.Create(group, []string{"contact-1"})

		assert.Error(t, err)
	})
}

func TestGroupService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		g1 := testutil.NewGroupFixture()
		g2 := testutil.NewGroupFixture()
		g2.ID = "group-2"
		mockRepo.EXPECT().FindAll("user-1").Return([]*domains.Group{g1, g2}, nil)

		result, err := svc.FindAll("user-1")

		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindAll("user-1").Return(nil, errors.New("db error"))

		_, err := svc.FindAll("user-1")
		assert.Error(t, err)
	})
}

func TestGroupService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)

		result, err := svc.FindByID("group-1", "user-1")

		assert.NoError(t, err)
		assert.Equal(t, "group-1", result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(nil, domains.NewNotFoundError("group not found"))

		_, err := svc.FindByID("group-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)

		_, err := svc.FindByID("group-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestGroupService_Update(t *testing.T) {
	t.Run("success - updates name only", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		existing := testutil.NewGroupFixture()
		incoming := testutil.NewGroupFixture()
		incoming.Name = "Viagem Europa"
		incoming.Category = ""

		mockRepo.EXPECT().FindByID("group-1").Return(existing, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(g *domains.Group) (*domains.Group, error) {
			assert.Equal(t, "Viagem Europa", g.Name)
			assert.Equal(t, "Travel", g.Category)
			return g, nil
		})

		result, err := svc.Update(incoming, "user-1")

		assert.NoError(t, err)
		assert.Equal(t, "Viagem Europa", result.Name)
	})

	t.Run("success - updates category only", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		existing := testutil.NewGroupFixture()
		incoming := testutil.NewGroupFixture()
		incoming.Name = ""
		incoming.Category = "Dinner"

		mockRepo.EXPECT().FindByID("group-1").Return(existing, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(g *domains.Group) (*domains.Group, error) {
			assert.Equal(t, "Viagem", g.Name)
			assert.Equal(t, "Dinner", g.Category)
			return g, nil
		})

		result, err := svc.Update(incoming, "user-1")

		assert.NoError(t, err)
		assert.Equal(t, "Dinner", result.Category)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(nil, domains.NewNotFoundError("group not found"))

		_, err := svc.Update(testutil.NewGroupFixture(), "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)

		_, err := svc.Update(testutil.NewGroupFixture(), "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestGroupService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)
		mockRepo.EXPECT().Delete("group-1").Return(nil)

		err := svc.Delete("group-1", "user-1")
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(nil, domains.NewNotFoundError("group not found"))

		err := svc.Delete("group-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)

		err := svc.Delete("group-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestGroupService_AddMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)
		mockRepo.EXPECT().IsContactOwnedBy("contact-1", "user-1").Return(true, nil)
		mockRepo.EXPECT().MemberExists("group-1", "contact-1").Return(false, nil)
		mockRepo.EXPECT().AddMember("group-1", "contact-1").Return(nil)

		err := svc.AddMember("group-1", "contact-1", "user-1")
		assert.NoError(t, err)
	})

	t.Run("group not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(nil, domains.NewNotFoundError("group not found"))

		err := svc.AddMember("group-1", "contact-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - not group owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)

		err := svc.AddMember("group-1", "contact-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("forbidden - contact not owned by user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)
		mockRepo.EXPECT().IsContactOwnedBy("contact-1", "user-1").Return(false, nil)

		err := svc.AddMember("group-1", "contact-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("conflict - already a member", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)
		mockRepo.EXPECT().IsContactOwnedBy("contact-1", "user-1").Return(true, nil)
		mockRepo.EXPECT().MemberExists("group-1", "contact-1").Return(true, nil)

		err := svc.AddMember("group-1", "contact-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrConflict)
	})
}

func TestGroupService_RemoveMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)
		mockRepo.EXPECT().MemberExists("group-1", "contact-1").Return(true, nil)
		mockRepo.EXPECT().RemoveMember("group-1", "contact-1").Return(nil)

		err := svc.RemoveMember("group-1", "contact-1", "user-1")
		assert.NoError(t, err)
	})

	t.Run("group not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(nil, domains.NewNotFoundError("group not found"))

		err := svc.RemoveMember("group-1", "contact-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - not group owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)

		err := svc.RemoveMember("group-1", "contact-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("member not found in group", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID("group-1").Return(testutil.NewGroupFixture(), nil)
		mockRepo.EXPECT().MemberExists("group-1", "contact-1").Return(false, nil)

		err := svc.RemoveMember("group-1", "contact-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}
