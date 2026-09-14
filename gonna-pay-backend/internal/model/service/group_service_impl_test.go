package service_test

import (
	"errors"
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/mocks"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/service"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGroupService_Create(t *testing.T) {
	t.Run("success without members", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		group := domains.NewGroup(m.OwnerID, m.Name, m.Category)
		mockRepo.EXPECT().Create(group, []string{}).Return(m.Group, nil)

		result, err := svc.Create(group, []string{})

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("success with members", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		group := domains.NewGroup(m.OwnerID, m.Name, m.Category)
		memberIDs := []string{testutil.ContactID, "contact-2"}

		mockRepo.EXPECT().IsContactOwnedBy(testutil.ContactID, m.OwnerID).Return(true, nil)
		mockRepo.EXPECT().IsContactOwnedBy("contact-2", m.OwnerID).Return(true, nil)
		mockRepo.EXPECT().Create(group, memberIDs).Return(m.Group, nil)

		result, err := svc.Create(group, memberIDs)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("forbidden - contact not owned by user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		group := domains.NewGroup(m.OwnerID, m.Name, m.Category)
		mockRepo.EXPECT().IsContactOwnedBy(testutil.ContactID, m.OwnerID).Return(false, nil)

		_, err := svc.Create(group, []string{testutil.ContactID})

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("repo error on IsContactOwnedBy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		group := domains.NewGroup(m.OwnerID, m.Name, m.Category)
		mockRepo.EXPECT().IsContactOwnedBy(testutil.ContactID, m.OwnerID).Return(false, errors.New("db error"))

		_, err := svc.Create(group, []string{testutil.ContactID})

		assert.Error(t, err)
	})
}

func TestGroupService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m1 := testutil.NewGroupMock()
		m2 := testutil.NewGroupMock()
		m2.Group.ID = "group-2"
		mockRepo.EXPECT().FindAll(testutil.UserID, 20, 0).Return([]*domains.Group{m1.Group, m2.Group}, int64(2), nil)

		result, total, err := svc.FindAll(testutil.UserID, 1, 20)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, int64(2), total)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindAll(testutil.UserID, 20, 0).Return(nil, int64(0), errors.New("db error"))

		_, _, err := svc.FindAll(testutil.UserID, 1, 20)
		assert.Error(t, err)
	})
}

func TestGroupService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)

		result, err := svc.FindByID(m.ID, testutil.UserID)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.GroupID).Return(nil, domains.NewNotFoundError("group not found"))

		_, err := svc.FindByID(testutil.GroupID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)

		_, err := svc.FindByID(m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestGroupService_Update(t *testing.T) {
	t.Run("success - updates name only", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		incoming := testutil.NewGroupFixture()
		incoming.Name = "Viagem Europa"
		incoming.Category = ""

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(g *domains.Group) (*domains.Group, error) {
			assert.Equal(t, "Viagem Europa", g.Name)
			assert.Equal(t, m.Category, g.Category)
			return g, nil
		})

		result, err := svc.Update(incoming, testutil.UserID)

		assert.NoError(t, err)
		assert.Equal(t, "Viagem Europa", result.Name)
	})

	t.Run("success - updates category only", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		incoming := testutil.NewGroupFixture()
		incoming.Name = ""
		incoming.Category = "Dinner"

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(g *domains.Group) (*domains.Group, error) {
			assert.Equal(t, m.Name, g.Name)
			assert.Equal(t, "Dinner", g.Category)
			return g, nil
		})

		result, err := svc.Update(incoming, testutil.UserID)

		assert.NoError(t, err)
		assert.Equal(t, "Dinner", result.Category)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.GroupID).Return(nil, domains.NewNotFoundError("group not found"))

		_, err := svc.Update(testutil.NewGroupFixture(), testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)

		_, err := svc.Update(testutil.NewGroupFixture(), testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestGroupService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().Delete(m.ID).Return(nil)

		err := svc.Delete(m.ID, testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.GroupID).Return(nil, domains.NewNotFoundError("group not found"))

		err := svc.Delete(testutil.GroupID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)

		err := svc.Delete(m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestGroupService_AddMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().IsContactOwnedBy(testutil.ContactID, testutil.UserID).Return(true, nil)
		mockRepo.EXPECT().MemberExists(m.ID, testutil.ContactID).Return(false, nil)
		mockRepo.EXPECT().AddMember(m.ID, testutil.ContactID).Return(nil)

		err := svc.AddMember(m.ID, testutil.ContactID, testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("group not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.GroupID).Return(nil, domains.NewNotFoundError("group not found"))

		err := svc.AddMember(testutil.GroupID, testutil.ContactID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - not group owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)

		err := svc.AddMember(m.ID, testutil.ContactID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("forbidden - contact not owned by user", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().IsContactOwnedBy(testutil.ContactID, testutil.UserID).Return(false, nil)

		err := svc.AddMember(m.ID, testutil.ContactID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("conflict - already a member", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().IsContactOwnedBy(testutil.ContactID, testutil.UserID).Return(true, nil)
		mockRepo.EXPECT().MemberExists(m.ID, testutil.ContactID).Return(true, nil)

		err := svc.AddMember(m.ID, testutil.ContactID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrConflict)
	})
}

func TestGroupService_RemoveMember(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().MemberExists(m.ID, testutil.ContactID).Return(true, nil)
		mockRepo.EXPECT().RemoveMember(m.ID, testutil.ContactID).Return(nil)

		err := svc.RemoveMember(m.ID, testutil.ContactID, testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("group not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.GroupID).Return(nil, domains.NewNotFoundError("group not found"))

		err := svc.RemoveMember(testutil.GroupID, testutil.ContactID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - not group owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)

		err := svc.RemoveMember(m.ID, testutil.ContactID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})

	t.Run("member not found in group", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockGroupRepository(ctrl)
		svc := service.NewGroupService(mockRepo)

		m := testutil.NewGroupMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Group, nil)
		mockRepo.EXPECT().MemberExists(m.ID, testutil.ContactID).Return(false, nil)

		err := svc.RemoveMember(m.ID, testutil.ContactID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}
