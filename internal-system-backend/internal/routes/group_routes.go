package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/auth"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/controller"
)

func GroupRoutes(router *gin.Engine, ctrl *controller.GroupController) {
	group := router.Group("/group")
	group.Use(auth.Middleware())
	{
		group.POST("", ctrl.CreateGroup)
		group.GET("/:id", ctrl.FindGroupByID)
		group.DELETE("/:id", ctrl.DeleteGroup)
		group.POST("/:id/member/:contactId", ctrl.AddMember)
		group.PATCH("/:id", ctrl.UpdateGroup)
		group.DELETE("/:id/member/:contactId", ctrl.RemoveMember)
	}

	groups := router.Group("/groups")
	groups.Use(auth.Middleware())
	{
		groups.GET("", ctrl.FindAllGroups)
	}
}
