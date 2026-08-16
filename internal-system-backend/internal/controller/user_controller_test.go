package controller_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/controller"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/mocks"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/testutil"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/response"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		user := testutil.NewUserFixture()
		mockService.EXPECT().Create(gomock.Any()).Return(user, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Rafael",
			"email":    "rafael@email.com",
			"password": "senha123",
			"phone":    "11999999999",
		})

		uc.CreateUser(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Rafael", body.Name)
		assert.Equal(t, "rafael@email.com", body.Email)
	})

	t.Run("validation error - missing fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name": "Rafael",
		})

		uc.CreateUser(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("validation error - invalid email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Rafael",
			"email":    "not-an-email",
			"password": "senha123",
			"phone":    "11999999999",
		})

		uc.CreateUser(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("service error - conflict", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		mockService.EXPECT().Create(gomock.Any()).Return(nil, domains.NewConflictError("email already in use"))

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     "Rafael",
			"email":    "rafael@email.com",
			"password": "senha123",
			"phone":    "11999999999",
		})

		uc.CreateUser(ctx)

		assert.Equal(t, http.StatusConflict, rec.Code)
	})
}

func TestFindAllUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		u1 := testutil.NewUserFixture()
		u2 := testutil.NewUserFixture()
		u2.ID = "user-2"
		mockService.EXPECT().FindAll().Return([]*domains.User{u1, u2}, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, nil)

		uc.FindAllUsers(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body []response.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Len(t, body, 2)
	})
}

func TestFindUserByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		user := testutil.NewUserFixture()
		mockService.EXPECT().FindByID("user-1").Return(user, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "user-1"}}, nil)

		uc.FindUserByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "user-1", body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		mockService.EXPECT().FindByID("user-1").Return(nil, domains.NewNotFoundError("user not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: "user-1"}}, nil)

		uc.FindUserByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		updated := testutil.NewUserFixture()
		updated.Name = "Rafael Novo"
		mockService.EXPECT().Update(gomock.Any()).Return(updated, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: "user-1"}}, map[string]any{
			"name": "Rafael Novo",
		})

		uc.UpdateUser(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Rafael Novo", body.Name)
	})

	t.Run("validation error - invalid email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: "user-1"}}, map[string]any{
			"email": "not-an-email",
		})

		uc.UpdateUser(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		mockService.EXPECT().Delete("user-1").Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "user-1"}})

		uc.DeleteUser(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		mockService.EXPECT().Delete("user-1").Return(domains.NewNotFoundError("user not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: "user-1"}})

		uc.DeleteUser(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
