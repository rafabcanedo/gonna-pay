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

func TestUserService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any()).Return(m.User, nil)

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		result, err := svc.Create(user)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("email already in use", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(m.Email).Return(m.User, nil)

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		_, err := svc.Create(user)

		assert.ErrorIs(t, err, domains.ErrConflict)
	})

	t.Run("password is encrypted before saving", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(u *domains.User) (*domains.User, error) {
			assert.NotEqual(t, "senha123", u.Password)
			return m.User, nil
		})

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		_, err := svc.Create(user)
		assert.NoError(t, err)
	})

	t.Run("repo error on FindByEmail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByEmail(testutil.UserEmail).Return(nil, errors.New("db error"))

		user := domains.NewUser(testutil.UserName, testutil.UserEmail, "senha123", testutil.UserPhone)
		_, err := svc.Create(user)

		assert.Error(t, err)
	})
}

func TestUserService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m1 := testutil.NewUserMock()
		m2 := testutil.NewUserMock()
		m2.User.ID = testutil.UserID2
		mockRepo.EXPECT().FindAll().Return([]*domains.User{m1.User, m2.User}, nil)

		result, err := svc.FindAll()

		assert.NoError(t, err)
		assert.Len(t, result, 2)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindAll().Return(nil, errors.New("db error"))

		_, err := svc.FindAll()
		assert.Error(t, err)
	})
}

func TestUserService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByID(m.ID).Return(m.User, nil)

		result, err := svc.FindByID(m.ID)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByID(testutil.UserID).Return(nil, domains.NewNotFoundError("user not found"))

		_, err := svc.FindByID(testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}

func TestUserService_FindByEmail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(m.Email).Return(m.User, nil)

		result, err := svc.FindByEmail(m.Email)

		assert.NoError(t, err)
		assert.Equal(t, m.Email, result.Email)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByEmail(testutil.UserEmail).Return(nil, domains.NewNotFoundError("not found"))

		_, err := svc.FindByEmail(testutil.UserEmail)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("success without password change", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		m.User.Name = testutil.UserUpdatedName
		mockRepo.EXPECT().FindByID(m.ID).Return(m.User, nil)
		mockRepo.EXPECT().FindByEmail(m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Update(m.User).Return(m.User, nil)

		result, err := svc.Update(m.User)

		assert.NoError(t, err)
		assert.Equal(t, testutil.UserUpdatedName, result.Name)
	})

	t.Run("success encrypts new password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		m.User.Password = "nova-senha"

		mockRepo.EXPECT().FindByID(m.ID).Return(m.User, nil)
		mockRepo.EXPECT().FindByEmail(m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(u *domains.User) (*domains.User, error) {
			assert.NotEqual(t, "nova-senha", u.Password)
			return m.User, nil
		})

		_, err := svc.Update(m.User)
		assert.NoError(t, err)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByID(testutil.UserID).Return(m.User, nil)
		mockRepo.EXPECT().Delete(testutil.UserID).Return(nil)

		err := svc.Delete(testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByID(testutil.UserID).Return(m.User, nil)
		mockRepo.EXPECT().Delete(testutil.UserID).Return(errors.New("db error"))

		err := svc.Delete(testutil.UserID)
		assert.Error(t, err)
	})
}
