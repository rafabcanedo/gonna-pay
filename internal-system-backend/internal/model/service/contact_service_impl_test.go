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

func TestContactService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		contact := testutil.NewContactFixture()
		mockRepo.EXPECT().Create(gomock.Any()).Return(contact, nil)

		result, err := svc.Create(domains.NewContact("user-1", "Ana", "ana@email.com", "11999999999", "Friend"))

		assert.NoError(t, err)
		assert.Equal(t, "contact-1", result.ID)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("db error"))

		_, err := svc.Create(domains.NewContact("user-1", "Ana", "ana@email.com", "11999999999", "Friend"))
		assert.Error(t, err)
	})
}

func TestContactService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		c1 := testutil.NewContactFixture()
		c2 := testutil.NewContactFixture()
		c2.ID = "contact-2"
		mockRepo.EXPECT().FindAll("user-1").Return([]*domains.Contact{c1, c2}, nil)

		result, err := svc.FindAll("user-1")

		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})
}

func TestContactService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID("contact-1").Return(testutil.NewContactFixture(), nil)

		result, err := svc.FindByID("contact-1", "user-1")

		assert.NoError(t, err)
		assert.Equal(t, "contact-1", result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID("contact-1").Return(nil, domains.NewNotFoundError("contact not found"))

		_, err := svc.FindByID("contact-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID("contact-1").Return(testutil.NewContactFixture(), nil)

		_, err := svc.FindByID("contact-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestContactService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		existing := testutil.NewContactFixture()
		incoming := domains.NewContactWithID("contact-1", "user-1", "Ana Silva", "nova@email.com", "11888888888", "Work")

		mockRepo.EXPECT().FindByID("contact-1").Return(existing, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(c *domains.Contact) (*domains.Contact, error) {
			assert.Equal(t, "Ana Silva", c.Name)
			assert.Equal(t, "nova@email.com", c.Email)
			assert.Equal(t, "11888888888", c.Phone)
			assert.Equal(t, "Work", c.Category)
			return c, nil
		})

		result, err := svc.Update(incoming)

		assert.NoError(t, err)
		assert.Equal(t, "Ana Silva", result.Name)
	})

	t.Run("success - partial update preserves existing fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		existing := testutil.NewContactFixture()
		incoming := domains.NewContactWithID("contact-1", "user-1", "Ana Silva", "", "", "")

		mockRepo.EXPECT().FindByID("contact-1").Return(existing, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(c *domains.Contact) (*domains.Contact, error) {
			assert.Equal(t, "Ana Silva", c.Name)
			assert.Equal(t, existing.Email, c.Email)
			assert.Equal(t, existing.Phone, c.Phone)
			assert.Equal(t, existing.Category, c.Category)
			return c, nil
		})

		_, err := svc.Update(incoming)
		assert.NoError(t, err)
	})

	t.Run("repo error on Update", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		existing := testutil.NewContactFixture()
		incoming := domains.NewContactWithID("contact-1", "user-1", "Ana Silva", "", "", "")

		mockRepo.EXPECT().FindByID("contact-1").Return(existing, nil)
		mockRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("db error"))

		_, err := svc.Update(incoming)
		assert.Error(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID("contact-1").Return(nil, domains.NewNotFoundError("not found"))

		incoming := testutil.NewContactFixture()
		_, err := svc.Update(incoming)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		existing := testutil.NewContactFixture()
		incoming := testutil.NewContactFixture()
		incoming.OwnerID = "user-2"

		mockRepo.EXPECT().FindByID("contact-1").Return(existing, nil)

		_, err := svc.Update(incoming)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestContactService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID("contact-1").Return(testutil.NewContactFixture(), nil)
		mockRepo.EXPECT().Delete("contact-1").Return(nil)

		err := svc.Delete("contact-1", "user-1")
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID("contact-1").Return(nil, domains.NewNotFoundError("not found"))

		err := svc.Delete("contact-1", "user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID("contact-1").Return(testutil.NewContactFixture(), nil)

		err := svc.Delete("contact-1", "user-2")

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}
