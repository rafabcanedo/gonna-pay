package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/auth"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/controller"
)

func AuthRoutes(router *gin.Engine, ctrl *controller.AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", ctrl.Login)
		authGroup.POST("/refresh", ctrl.Refresh)
		authGroup.POST("/logout", ctrl.Logout)

		authGroup.GET("/profile", auth.Middleware(), ctrl.GetProfile)
		authGroup.PUT("/profile", auth.Middleware(), ctrl.UpdateProfile)
		authGroup.DELETE("/profile", auth.Middleware(), ctrl.DeleteProfile)
	}
}
