package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/auth"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/controller"
)

func AuthRoutes(router *gin.Engine, ctrl *controller.AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", ctrl.Login)
		authGroup.POST("/refresh", ctrl.Refresh)
		authGroup.POST("/logout", ctrl.Logout)
		authGroup.POST("/verify-email", ctrl.VerifyEmail)
		authGroup.POST("/forgot-password", ctrl.ForgotPassword)
		authGroup.POST("/reset-password", ctrl.ResetPassword)

		authGroup.GET("/profile", auth.Middleware(), ctrl.GetProfile)
		authGroup.PUT("/profile", auth.Middleware(), ctrl.UpdateProfile)
		authGroup.DELETE("/profile", auth.Middleware(), ctrl.DeleteProfile)
	}
}
