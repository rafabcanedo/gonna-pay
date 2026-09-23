package controller_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/controller"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/mocks"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/repository/entity"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/testutil"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func init() {
	os.Setenv("JWT_SECRET", "test-secret-key")
}

func hashPassword(password string) (string, error) {
	user := domains.NewUser("", "", password, "")
	if err := user.EncryptPassword(); err != nil {
		return "", err
	}
	return user.Password, nil
}

func hasCookie(cookies []*http.Cookie, name string) bool {
	for _, c := range cookies {
		if c.Name == name {
			return true
		}
	}
	return false
}

func hasCookieExpired(cookies []*http.Cookie, name string) bool {
	for _, c := range cookies {
		if c.Name == name && c.MaxAge < 0 {
			return true
		}
	}
	return false
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		user := testutil.NewUserFixture()
		user.Password, _ = hashPassword("senha123")

		mockUserService.EXPECT().FindByEmail(gomock.Any(), "rafael@email.com").Return(user, nil)
		mockAuthRepo.EXPECT().Save(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    "rafael@email.com",
			"password": "senha123",
		})

		ac.Login(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body map[string]any
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Login successful", body["message"])
		assert.NotNil(t, body["user"])
		cookies := rec.Result().Cookies()
		assert.True(t, hasCookie(cookies, "access_token"))
		assert.True(t, hasCookie(cookies, "refresh_token"))
	})

	t.Run("validation error - invalid email", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    "not-an-email",
			"password": "senha123",
		})

		ac.Login(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("user not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		mockUserService.EXPECT().FindByEmail(gomock.Any(), "rafael@email.com").Return(nil, domains.NewNotFoundError("not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    "rafael@email.com",
			"password": "senha123",
		})

		ac.Login(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("wrong password", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		user := testutil.NewUserFixture()
		user.Password, _ = hashPassword("senha-correta")
		mockUserService.EXPECT().FindByEmail(gomock.Any(), "rafael@email.com").Return(user, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, map[string]any{
			"email":    "rafael@email.com",
			"password": "senha-errada",
		})

		ac.Login(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}

func TestRefresh(t *testing.T) {
	t.Run("missing cookie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, nil)

		ac.Refresh(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("invalid token - not found in db", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		mockAuthRepo.EXPECT().FindByHash(gomock.Any(), gomock.Any()).Return(nil, domains.NewNotFoundError("not found"))

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, nil)
		ctx.Request.Header.Set("Cookie", "refresh_token=sometoken")

		ac.Refresh(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("expired token", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		expiredToken := &entity.RefreshTokenEntity{
			ID:        uuid.New(),
			UserID:    uuid.New(),
			TokenHash: "somehash",
			ExpiresAt: time.Now().Add(-1 * time.Hour),
		}
		mockAuthRepo.EXPECT().FindByHash(gomock.Any(), gomock.Any()).Return(expiredToken, nil)
		mockAuthRepo.EXPECT().DeleteByHash(gomock.Any(), gomock.Any()).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, nil)
		ctx.Request.Header.Set("Cookie", "refresh_token=sometoken")

		ac.Refresh(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}

func TestLogout(t *testing.T) {
	t.Run("success with cookie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		mockAuthRepo.EXPECT().DeleteByHash(gomock.Any(), gomock.Any()).Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, nil)
		ctx.Request.Header.Set("Cookie", "refresh_token=sometoken")

		ac.Logout(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		cookies := rec.Result().Cookies()
		assert.True(t, hasCookieExpired(cookies, "access_token"))
		assert.True(t, hasCookieExpired(cookies, "refresh_token"))
	})

	t.Run("success without cookie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		ctx, rec := testutil.NewTestContext()
		testutil.MakePost(ctx, nil, nil)

		ac.Logout(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}

func TestGetProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		user := testutil.NewUserFixture()
		mockUserService.EXPECT().FindByID(gomock.Any(), "user-1").Return(user, nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeGet(ctx, nil, nil)

		ac.GetProfile(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		var body map[string]any
		json.Unmarshal(rec.Body.Bytes(), &body)
		assert.Equal(t, "Rafael", body["name"])
	})

	t.Run("missing userID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeGet(ctx, nil, nil)

		ac.GetProfile(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}

func TestDeleteProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		mockUserService.EXPECT().Delete(gomock.Any(), "user-1").Return(nil)

		ctx, rec := testutil.NewTestContext()
		testutil.SetAuthUser(ctx, "user-1")
		testutil.MakeDelete(ctx, nil)

		ac.DeleteProfile(ctx)

		assert.Equal(t, http.StatusOK, rec.Code)
		cookies := rec.Result().Cookies()
		assert.True(t, hasCookieExpired(cookies, "access_token"))
	})

	t.Run("missing userID", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockUserService := mocks.NewMockUserService(ctrl)
		mockAuthRepo := mocks.NewMockAuthRepository(ctrl)
		ac := controller.NewAuthController(mockUserService, mockAuthRepo)

		ctx, rec := testutil.NewTestContext()
		testutil.MakeDelete(ctx, nil)

		ac.DeleteProfile(ctx)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})
}
