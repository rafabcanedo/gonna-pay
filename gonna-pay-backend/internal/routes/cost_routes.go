package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/auth"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/controller"
)

func CostRoutes(router *gin.Engine, ctrl *controller.CostController) {
	cost := router.Group("/cost")
	cost.Use(auth.Middleware())
	{
		cost.POST("", ctrl.CreateCost)
		cost.GET("/:id", ctrl.FindCostByID)
		cost.PATCH("/:id", ctrl.UpdateCost)
		cost.DELETE("/:id", ctrl.DeleteCost)
	}

	costs := router.Group("/costs")
	costs.Use(auth.Middleware())
	{
		costs.GET("", ctrl.FindAllCosts)
	}
}
