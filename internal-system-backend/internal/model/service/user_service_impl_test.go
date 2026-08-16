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

func TestUserService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		created := testutil.NewUserFixture()
		mockRepo.EXPECT().FindByEmail("rafael@email.com").Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any()).Return(created, nil)

		user := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")
		result, err := svc.Create(user)

		assert.NoError(t, err)
		assert.Equal(t, "user-1", result.ID)
	})

	t.Run("email already in use", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByEmail("rafael@email.com").Return(testutil.NewUserFixture(), nil)

		user := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")
		_, err := svc.Create(user)

		assert.ErrorIs(t, err, domains.ErrConflict)
	})

	t.Run("password is encrypted before saving", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByEmail("rafael@email.com").Return(nil, domains.NewNotFoundError("not found"))
		mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(u *domains.User) (*domains.User, error) {
			assert.NotEqual(t, "senha123", u.Password)
			return testutil.NewUserFixture(), nil
		})

		user := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")
		_, err := svc.Create(user)
		assert.NoError(t, err)
	})

	t.Run("repo error on FindByEmail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByEmail("rafael@email.com").Return(nil, errors.New("db error"))

		user := domains.NewUser("Rafael", "rafael@email.com", "senha123", "11999999999")
		_, err := svc.Create(user)

		assert.Error(t, err)
	})
}

func TestUserService_FindAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		u1 := testutil.NewUserFixture()
		u2 := testutil.NewUserFixture()
		u2.ID = "user-2"
		mockRepo.EXPECT().FindAll().Return([]*domains.User{u1, u2}, nil)

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

		mockRepo.EXPECT().FindByID("user-1").Return(testutil.NewUserFixture(), nil)

		result, err := svc.FindByID("user-1")

		assert.NoError(t, err)
		assert.Equal(t, "user-1", result.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByID("user-1").Return(nil, domains.NewNotFoundError("user not found"))

		_, err := svc.FindByID("user-1")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}

func TestUserService_FindByEmail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByEmail("rafael@email.com").Return(testutil.NewUserFixture(), nil)

		result, err := svc.FindByEmail("rafael@email.com")

		assert.NoError(t, err)
		assert.Equal(t, "rafael@email.com", result.Email)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().FindByEmail("rafael@email.com").Return(nil, domains.NewNotFoundError("not found"))

		_, err := svc.FindByEmail("rafael@email.com")

		assert.ErrorIs(t, err, domains.ErrNotFound)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("success without password change", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		user := testutil.NewUserFixture()
		user.Name = "Rafael Novo"
		mockRepo.EXPECT().Update(user).Return(user, nil)

		result, err := svc.Update(user)

		assert.NoError(t, err)
		assert.Equal(t, "Rafael Novo", result.Name)
	})

	t.Run("success encrypts new password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		user := testutil.NewUserFixture()
		user.Password = "nova-senha"

		mockRepo.EXPECT().Update(gomock.Any()).DoAndReturn(func(u *domains.User) (*domains.User, error) {
			assert.NotEqual(t, "nova-senha", u.Password)
			return testutil.NewUserFixture(), nil
		})

		_, err := svc.Update(user)
		assert.NoError(t, err)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().Delete("user-1").Return(nil)

		err := svc.Delete("user-1")
		assert.NoError(t, err)
	})

	t.Run("repo error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockUserRepository(ctrl)
		svc := service.NewUserService(mockRepo)

		mockRepo.EXPECT().Delete("user-1").Return(errors.New("db error"))

		err := svc.Delete("user-1")
		assert.Error(t, err)
	})
}
