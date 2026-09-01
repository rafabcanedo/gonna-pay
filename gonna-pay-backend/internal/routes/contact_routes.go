package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/auth"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/controller"
)

func ContactRoutes(router *gin.Engine, ctrl *controller.ContactController) {
	contact := router.Group("/contact")
    contact.Use(auth.Middleware())
    {
        contact.POST("", ctrl.CreateContact)
        contact.GET("/:id", ctrl.FindContactByID)
        contact.PATCH("/:id", ctrl.UpdateContact)
        contact.DELETE("/:id", ctrl.DeleteContact)
    }

    contacts := router.Group("/contacts")
    contacts.Use(auth.Middleware())
    {
        contacts.GET("", ctrl.FindAllContacts)
    }
}
