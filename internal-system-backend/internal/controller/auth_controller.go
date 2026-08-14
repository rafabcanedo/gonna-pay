package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/auth"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/logger"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/rest_errors"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/validation"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/repository"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/service"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/request"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/response"
)

type AuthController struct {
	userService service.UserService
	authRepo    repository.AuthRepository
}

func NewAuthController(userService service.UserService, authRepo repository.AuthRepository) *AuthController {
	return &AuthController{userService: userService, authRepo: authRepo}
}

func (ac *AuthController) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	user, restErr := ac.userService.FindByEmail(req.Email)
	if restErr != nil {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("invalid credentials"))
		return
	}

	if err := user.ComparePassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("invalid credentials"))
		return
	}

	accessToken, err := auth.GenerateAccessToken(user.GetID(), user.GetName())
	if err != nil {
		logger.Error("error generating access token", err)
		restErr := rest_errors.NewInternalServerError("error generating token")
		c.JSON(restErr.Code, restErr)
		return
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		logger.Error("error generating refresh token", err)
		restErr := rest_errors.NewInternalServerError("error generating token")
		c.JSON(restErr.Code, restErr)
		return
	}

	tokenHash := auth.HashToken(refreshToken)
	expiresAt := time.Now().Add(48 * time.Hour)

	if err := ac.authRepo.Save(user.GetID(), tokenHash, expiresAt); err != nil {
		logger.Error("error saving refresh token", err)
		restErr := rest_errors.NewInternalServerError("error processing login")
		c.JSON(restErr.Code, restErr)
		return
	}

	c.SetCookie("access_token", accessToken, 1800, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 172800, "/auth", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": response.UserResponse{
			ID:    user.GetID(),
			Name:  user.GetName(),
			Email: user.GetEmail(),
			Phone: user.GetPhone(),
		},
	})
}

func (ac *AuthController) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		restErr := rest_errors.NewUnauthorizedRequestError("missing refresh token")
		c.JSON(restErr.Code, restErr)
		return
	}

	tokenHash := auth.HashToken(refreshToken)

	stored, err := ac.authRepo.FindByHash(tokenHash)
	if err != nil {
		restErr := rest_errors.NewUnauthorizedRequestError("invalid refresh token")
		c.JSON(restErr.Code, restErr)
		return
	}

	if time.Now().After(stored.ExpiresAt) {
		ac.authRepo.DeleteByHash(tokenHash)
		restErr := rest_errors.NewUnauthorizedRequestError("refresh token expired")
		c.JSON(restErr.Code, restErr)
		return
	}

	ac.authRepo.DeleteByHash(tokenHash)

	user, restErr := ac.userService.FindByID(stored.UserID.String())
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	newAccessToken, err := auth.GenerateAccessToken(user.GetID(), user.GetName())
	if err != nil {
		logger.Error("error generating access token on refresh", err)
		restErr := rest_errors.NewInternalServerError("error generating token")
		c.JSON(restErr.Code, restErr)
		return
	}

	newRefreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		logger.Error("error generating refresh token on refresh", err)
		restErr := rest_errors.NewInternalServerError("error generating token")
		c.JSON(restErr.Code, restErr)
		return
	}

	newTokenHash := auth.HashToken(newRefreshToken)
	newExpiresAt := time.Now().Add(48 * time.Hour)

	if err := ac.authRepo.Save(user.GetID(), newTokenHash, newExpiresAt); err != nil {
		logger.Error("error saving new refresh token", err)
		restErr := rest_errors.NewInternalServerError("error processing refresh")
		c.JSON(restErr.Code, restErr)
		return
	}

	c.SetCookie("access_token", newAccessToken, 1800, "/", "", false, true)
	c.SetCookie("refresh_token", newRefreshToken, 172800, "/auth", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed"})
}

func (ac *AuthController) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err == nil {
		tokenHash := auth.HashToken(refreshToken)
		ac.authRepo.DeleteByHash(tokenHash)
	}

	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/auth", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (ac *AuthController) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		restErr := rest_errors.NewUnauthorizedRequestError("user identification missing")
		c.JSON(restErr.Code, restErr)
		return
	}

	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	current, restErr := ac.userService.FindByID(userID)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	if req.Name != "" {
		current.SetName(req.Name)
	}
	if req.Email != "" {
		current.SetEmail(req.Email)
	}
	if req.Phone != "" {
		current.SetPhone(req.Phone)
	}

	updated, restErr := ac.userService.Update(current)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": response.UserResponse{
			ID:    updated.GetID(),
			Name:  updated.GetName(),
			Email: updated.GetEmail(),
			Phone: updated.GetPhone(),
		},
	})
}

func (ac *AuthController) GetProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		restErr := rest_errors.NewUnauthorizedRequestError("user identification missing")
		c.JSON(restErr.Code, restErr)
		return
	}

	user, restErr := ac.userService.FindByID(userID)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response.UserResponse{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
		Phone: user.GetPhone(),
	})
}
