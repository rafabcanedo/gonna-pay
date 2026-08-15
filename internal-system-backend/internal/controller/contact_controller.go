package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/validation"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/service"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/request"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/view/response"
)

type ContactController struct {
	service service.ContactService
}

func NewContactController(service service.ContactService) *ContactController {
	return &ContactController{service: service}
}

func (cc *ContactController) CreateContact(c *gin.Context) {
	ownerID := c.GetString("userID")

	var req request.CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domains.NewContactDomain(ownerID, req.Name, req.Email, req.Phone, req.Category)

	created, err := cc.service.Create(domain)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewContactResponse(created))
}

func (cc *ContactController) FindAllContacts(c *gin.Context) {
	ownerID := c.GetString("userID")

	contacts, err := cc.service.FindAll(ownerID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewContactResponseList(contacts))
}

func (cc *ContactController) FindContactByID(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	contact, err := cc.service.FindByID(id, ownerID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewContactResponse(contact))
}

func (cc *ContactController) UpdateContact(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	var req request.UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domains.NewContactDomainWithID(id, ownerID, req.Name, req.Email, req.Phone, req.Category)

	updated, err := cc.service.Update(domain)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewContactResponse(updated))
}

func (cc *ContactController) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	if err := cc.service.Delete(id, ownerID); err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contact deleted successfully"})
}
