package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/mocks"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/service"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func newUserSvc(ctrl *gomock.Controller) (service.UserService, *mocks.MockUserRepository, *mocks.MockEmailTokenRepository, *mocks.MockEmailService) {
	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockEmailTokenRepo := mocks.NewMockEmailTokenRepository(ctrl)
	mockEmailSvc := mocks.NewMockEmailService(ctrl)
	svc := service.NewUserService(mockRepo, mockEmailTokenRepo, mockEmailSvc)
	return svc, mockRepo, mockEmailTokenRepo, mockEmailSvc
}

func TestUserService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, mockEmailTokenRepo, mockEmailSvc := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(m.User, nil)
		mockEmailTokenRepo.EXPECT().Save(gomock.Any(), m.ID, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockEmailSvc.EXPECT().SendVerificationEmail(gomock.Any(), m.Email, m.Name, gomock.Any()).Return(nil)

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		result, err := svc.Create(context.Background(), user)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("email already in use", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(m.User, nil)

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		_, err := svc.Create(context.Background(), user)

		assert.ErrorIs(t, err, domains.ErrConflict)
	})

	t.Run("password is encrypted before saving", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, mockEmailTokenRepo, mockEmailSvc := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, u *domains.User) (*domains.User, error) {
			assert.NotEqual(t, "senha123", u.Password)
			return m.User, nil
		})
		mockEmailTokenRepo.EXPECT().Save(gomock.Any(), m.ID, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockEmailSvc.EXPECT().SendVerificationEmail(gomock.Any(), m.Email, m.Name, gomock.Any()).Return(nil)

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		_, err := svc.Create(context.Background(), user)
		assert.NoError(t, err)
	})

	t.Run("repo error on FindByEmail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		mockRepo.EXPECT().FindByEmail(gomock.Any(), testutil.UserEmail).Return(nil, errors.New("db error"))

		user := domains.NewUser(testutil.UserName, testutil.UserEmail, "senha123", testutil.UserPhone)
		_, err := svc.Create(context.Background(), user)

		assert.Error(t, err)
	})

	t.Run("email token save error does not fail create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, mockEmailTokenRepo, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(m.User, nil)
		mockEmailTokenRepo.EXPECT().Save(gomock.Any(), m.ID, gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("db error"))

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		result, err := svc.Create(context.Background(), user)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("email send error does not fail create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, mockEmailTokenRepo, mockEmailSvc := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(m.User, nil)
		mockEmailTokenRepo.EXPECT().Save(gomock.Any(), m.ID, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockEmailSvc.EXPECT().SendVerificationEmail(gomock.Any(), m.Email, m.Name, gomock.Any()).Return(errors.New("resend error"))

		user := domains.NewUser(m.Name, m.Email, "senha123", m.Phone)
		result, err := svc.Create(context.Background(), user)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})
}

func TestUserService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m1 := testutil.NewUserMock()
		m2 := testutil.NewUserMock()
		m2.User.ID = testutil.UserID2
		mockRepo.EXPECT().FindAll(gomock.Any(), 20, 0).Return([]*domains.User{m1.User, m2.User}, int64(2), nil)

		result, total, err := svc.FindAll(context.Background(), 1, 20)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, int64(2), total)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		mockRepo.EXPECT().FindAll(gomock.Any(), 20, 0).Return(nil, int64(0), errors.New("db error"))

		_, _, err := svc.FindAll(context.Background(), 1, 20)
		assert.Error(t, err)
	})
}

func TestUserService_FindByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.User, nil)

		result, err := svc.FindByID(context.Background(), m.ID)

		assert.NoError(t, err)
		assert.Equal(t, m.ID, result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		mockRepo.EXPECT().FindByID(gomock.Any(), testutil.UserID).Return(nil, domains.NewNotFoundError("user not found"))

		_, err := svc.FindByID(context.Background(), testutil.UserID)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}

func TestUserService_FindByEmail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(m.User, nil)

		result, err := svc.FindByEmail(context.Background(), m.Email)

		assert.NoError(t, err)
		assert.Equal(t, m.Email, result.Email)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		mockRepo.EXPECT().FindByEmail(gomock.Any(), testutil.UserEmail).Return(nil, domains.NewNotFoundError("not found"))

		_, err := svc.FindByEmail(context.Background(), testutil.UserEmail)

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("success without password change", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		m.User.Name = testutil.UserUpdatedName
		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.User, nil)
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Update(gomock.Any(), m.User).Return(m.User, nil)

		result, err := svc.Update(context.Background(), m.User)

		assert.NoError(t, err)
		assert.Equal(t, testutil.UserUpdatedName, result.Name)
	})

	t.Run("success encrypts new password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		m.User.Password = "nova-senha"

		mockRepo.EXPECT().FindByID(gomock.Any(), m.ID).Return(m.User, nil)
		mockRepo.EXPECT().FindByEmail(gomock.Any(), m.Email).Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, u *domains.User) (*domains.User, error) {
			assert.NotEqual(t, "nova-senha", u.Password)
			return m.User, nil
		})

		_, err := svc.Update(context.Background(), m.User)
		assert.NoError(t, err)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), testutil.UserID).Return(m.User, nil)
		mockRepo.EXPECT().Delete(gomock.Any(), testutil.UserID).Return(nil)

		err := svc.Delete(context.Background(), testutil.UserID)
		assert.NoError(t, err)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		svc, mockRepo, _, _ := newUserSvc(ctrl)

		m := testutil.NewUserMock()
		mockRepo.EXPECT().FindByID(gomock.Any(), testutil.UserID).Return(m.User, nil)
		mockRepo.EXPECT().Delete(gomock.Any(), testutil.UserID).Return(errors.New("db error"))

		err := svc.Delete(context.Background(), testutil.UserID)
		assert.Error(t, err)
	})
}
