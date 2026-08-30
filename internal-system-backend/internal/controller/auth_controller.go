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

// @Summary      Login
// @Description  Autentica o usuário e retorna cookies de sessão (access_token e refresh_token)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      request.LoginRequest  true  "Credenciais do usuário"
// @Success      200   {object}  map[string]interface{}  "message e user"
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      500   {object}  rest_errors.RestErrors
// @Router       /auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	user, err := ac.userService.FindByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("invalid credentials"))
		return
	}

	if err := user.ComparePassword(req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("invalid credentials"))
		return
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, user.Name)
	if err != nil {
		logger.Error("error generating access token", err)
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError("error generating token"))
		return
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		logger.Error("error generating refresh token", err)
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError("error generating token"))
		return
	}

	tokenHash := auth.HashToken(refreshToken)
	expiresAt := time.Now().Add(48 * time.Hour)

	if err := ac.authRepo.Save(user.ID, tokenHash, expiresAt); err != nil {
		logger.Error("error saving refresh token", err)
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError("error processing login"))
		return
	}

	c.SetCookie("access_token", accessToken, 1800, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 172800, "/auth", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    response.NewUserResponse(user),
	})
}

// @Summary      Refresh token
// @Description  Gera um novo access_token usando o cookie refresh_token. O refresh_token antigo é invalidado (rotação)
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Router       /auth/refresh [post]
func (ac *AuthController) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("missing refresh token"))
		return
	}

	tokenHash := auth.HashToken(refreshToken)

	stored, err := ac.authRepo.FindByHash(tokenHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("invalid refresh token"))
		return
	}

	if time.Now().After(stored.ExpiresAt) {
		ac.authRepo.DeleteByHash(tokenHash)
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("refresh token expired"))
		return
	}

	ac.authRepo.DeleteByHash(tokenHash)

	user, err := ac.userService.FindByID(stored.UserID.String())
	if err != nil {
		response.RespondError(c, err)
		return
	}

	newAccessToken, err := auth.GenerateAccessToken(user.ID, user.Name)
	if err != nil {
		logger.Error("error generating access token on refresh", err)
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError("error generating token"))
		return
	}

	newRefreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		logger.Error("error generating refresh token on refresh", err)
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError("error generating token"))
		return
	}

	newTokenHash := auth.HashToken(newRefreshToken)
	newExpiresAt := time.Now().Add(48 * time.Hour)

	if err := ac.authRepo.Save(user.ID, newTokenHash, newExpiresAt); err != nil {
		logger.Error("error saving new refresh token", err)
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError("error processing refresh"))
		return
	}

	c.SetCookie("access_token", newAccessToken, 1800, "/", "", false, true)
	c.SetCookie("refresh_token", newRefreshToken, 172800, "/auth", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed"})
}

// @Summary      Logout
// @Description  Invalida o refresh_token e limpa os cookies de sessão
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string  "message"
// @Router       /auth/logout [post]
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

// @Summary      Atualizar perfil
// @Description  Atualiza os dados do usuário autenticado (campos opcionais, apenas os enviados são alterados)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      request.UpdateProfileRequest  true  "Dados para atualização"
// @Success      200   {object}  map[string]interface{}  "message e user"
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /auth/profile [put]
func (ac *AuthController) UpdateProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("user identification missing"))
		return
	}

	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	current, err := ac.userService.FindByID(userID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	if req.Name != "" {
		current.Name = req.Name
	}
	if req.Email != "" {
		current.Email = req.Email
	}
	if req.Phone != "" {
		current.Phone = req.Phone
	}

	updated, err := ac.userService.Update(current)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user":    response.NewUserResponse(updated),
	})
}

// @Summary      Buscar perfil
// @Description  Retorna os dados do usuário autenticado
// @Tags         auth
// @Produce      json
// @Success      200  {object}  response.UserResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /auth/profile [get]
func (ac *AuthController) GetProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("user identification missing"))
		return
	}

	user, err := ac.userService.FindByID(userID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(user))
}

// @Summary      Deletar conta
// @Description  Remove a conta do usuário autenticado e limpa os cookies de sessão
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /auth/profile [delete]
func (ac *AuthController) DeleteProfile(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError("user identification missing"))
		return
	}

	if err := ac.userService.Delete(userID); err != nil {
		response.RespondError(c, err)
		return
	}

	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/auth", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "account deleted successfully"})
}
