package controller_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/controller"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/mocks"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/response"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		m := testutil.NewUserMock()
		mockService.EXPECT().Create(gomock.Any()).Return(m.User, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name":     m.Name,
			"email":    m.Email,
			"password": "senha123",
			"phone":    m.Phone,
		})

		uc.CreateUser(ctx)

		assert.Equal(t, http.StatusCreated, rec.Code)
		var body response.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.Name, body.Name)
		assert.Equal(t, m.Email, body.Email)
	})

	t.Run("validation error - missing fields", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"name": testutil.UserName,
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
			"name":     testutil.UserName,
			"email":    "not-an-email",
			"password": "senha123",
			"phone":    testutil.UserPhone,
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
			"name":     testutil.UserName,
			"email":    testutil.UserEmail,
			"password": "senha123",
			"phone":    testutil.UserPhone,
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

		m1 := testutil.NewUserMock()
		m2 := testutil.NewUserMock()
		m2.User.ID = testutil.UserID2
		mockService.EXPECT().FindAll(1, 20).Return([]*domains.User{m1.User, m2.User}, int64(2), nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, nil)

		uc.FindAllUsers(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.PaginatedResponse[response.UserResponse]
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, int64(2), body.Total)
		assert.Equal(t, 1, body.TotalPages)
	})

	t.Run("custom page and limit", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		m := testutil.NewUserMock()
		mockService.EXPECT().FindAll(2, 10).Return([]*domains.User{m.User}, int64(11), nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, url.Values{"page": {"2"}, "limit": {"10"}})

		uc.FindAllUsers(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.PaginatedResponse[response.UserResponse]
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, 2, body.Page)
		assert.Equal(t, 10, body.Limit)
		assert.Equal(t, int64(11), body.Total)
		assert.Equal(t, 2, body.TotalPages)
	})

	t.Run("service error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		mockService.EXPECT().FindAll(1, 20).Return(nil, int64(0), domains.NewNotFoundError("not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, nil)

		uc.FindAllUsers(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestFindUserByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		m := testutil.NewUserMock()
		mockService.EXPECT().FindByID(m.ID).Return(m.User, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: m.ID}}, nil)

		uc.FindUserByID(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, m.ID, body.ID)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		mockService.EXPECT().FindByID(testutil.UserID).Return(nil, domains.NewNotFoundError("user not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, gin.Params{{Key: "id", Value: testutil.UserID}}, nil)

		uc.FindUserByID(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		m := testutil.NewUserMock()
		m.User.Name = testutil.UserUpdatedName
		mockService.EXPECT().Update(gomock.Any()).Return(m.User, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, m.ID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: m.ID}}, map[string]any{
			"name": testutil.UserUpdatedName,
		})

		uc.UpdateUser(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body response.UserResponse
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, testutil.UserUpdatedName, body.Name)
	})

	t.Run("validation error - invalid email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakePatch(ctx, gin.Params{{Key: "id", Value: testutil.UserID}}, map[string]any{
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

		mockService.EXPECT().Delete(testutil.UserID).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.UserID}})

		uc.DeleteUser(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockService := mocks.NewMockUserService(ctrl)
		uc := controller.NewUserController(mockService)

		mockService.EXPECT().Delete(testutil.UserID).Return(domains.NewNotFoundError("user not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, testutil.UserID)
		testutil.MakeDelete(ctx, gin.Params{{Key: "id", Value: testutil.UserID}})

		uc.DeleteUser(ctx)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
