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

func TestContactService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		mockRepo.EXPECT().Create(gomock.Any()).Return(m.Contact, nil)

		result, err := svc.Create(domains.NewContact(m.OwnerID, m.Name, m.Email, m.Phone, m.Category))

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("db error"))

		_, err := svc.Create(domains.NewContact(testutil.UserID, testutil.ContactName, testutil.ContactEmail, testutil.ContactPhone, testutil.ContactCategory))
		assert.Error(t, err)
	})
}

func TestContactService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m1 := testutil.NewContactMock()
		m2 := testutil.NewContactMock()
		m2.Contact.ID = "contact-2"
		mockRepo.EXPECT().FindAll(testutil.UserID).Return([]*domains.Contact{m1.Contact, m2.Contact}, nil)

		result, err := svc.FindAll(testutil.UserID)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})
}

func TestContactService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)

		result, err := svc.FindByID(m.ID, testutil.UserID)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.ContactID).Return(nil, domains.NewNotFoundError("contact not found"))

		_, err := svc.FindByID(testutil.ContactID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)

		_, err := svc.FindByID(m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestContactService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		incoming := domains.NewContactWithID(m.ID, m.OwnerID, testutil.ContactUpdatedName, "nova@email.com", "11888888888", "Work")

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(c *domains.Contact) (*domains.Contact, error) {
			assert.Equal(t, testutil.ContactUpdatedName, c.Name)
			assert.Equal(t, "nova@email.com", c.Email)
			assert.Equal(t, "11888888888", c.Phone)
			assert.Equal(t, "Work", c.Category)
			return c, nil
		})

		result, err := svc.Update(incoming)

		assert.NoError(t, err)
		assert.Equal(t, testutil.ContactUpdatedName, result.Name)
	})

	t.Run("success - partial update preserves existing fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		incoming := domains.NewContactWithID(m.ID, m.OwnerID, testutil.ContactUpdatedName, "", "", "")

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(c *domains.Contact) (*domains.Contact, error) {
			assert.Equal(t, testutil.ContactUpdatedName, c.Name)
			assert.Equal(t, m.Email, c.Email)
			assert.Equal(t, m.Phone, c.Phone)
			assert.Equal(t, m.Category, c.Category)
			return c, nil
		})

		_, err := svc.Update(incoming)
		assert.NoError(t, err)
	})

	t.Run("repo error on Update", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		incoming := domains.NewContactWithID(m.ID, m.OwnerID, testutil.ContactUpdatedName, "", "", "")

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)
		mockRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("db error"))

		_, err := svc.Update(incoming)
		assert.Error(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.ContactID).Return(nil, domains.NewNotFoundError("not found"))

		incoming := testutil.NewContactFixture()
		_, err := svc.Update(incoming)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		incoming := testutil.NewContactFixture()
		incoming.OwnerID = testutil.UserID2

		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)

		_, err := svc.Update(incoming)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}

func TestContactService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)
		mockRepo.EXPECT().Delete(m.ID).Return(nil)

		err := svc.Delete(m.ID, testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.ContactID).Return(nil, domains.NewNotFoundError("not found"))

		err := svc.Delete(testutil.ContactID, testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})

	t.Run("forbidden - different owner", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockContactRepository(ctrl)
		svc := service.NewContactService(mockRepo)

		m := testutil.NewContactMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.Contact, nil)

		err := svc.Delete(m.ID, testutil.UserID2)

		assert.ErrorIs(t, err, domains.ErrForbidden)
	})
}
