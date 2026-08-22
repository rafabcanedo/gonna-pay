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

type CostController struct {
	service service.CostService
}

func NewCostController(service service.CostService) *CostController {
	return &CostController{service: service}
}

func (cc *CostController) CreateCost(c *gin.Context) {
	userID := c.GetString("userID")

	var req request.CreateCostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	cost := domains.NewCost(userID, req.GroupID, req.CostName, req.Category, req.TotalValue, 0)

	created, err := cc.service.Create(cost, req.OwnerPercentage)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewCostDetailResponse(created))
}

func (cc *CostController) FindAllCosts(c *gin.Context) {
	userID := c.GetString("userID")

	costs, err := cc.service.FindAll(userID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewCostResponseList(costs))
}

func (cc *CostController) FindCostByID(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	cost, err := cc.service.FindByID(id, userID)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewCostDetailResponse(cost))
}

func (cc *CostController) UpdateCost(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	var req request.UpdateCostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	cost := domains.NewCost(userID, "", req.CostName, req.Category, req.TotalValue, 0)

	updated, err := cc.service.Update(id, userID, cost, req.OwnerPercentage)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewCostDetailResponse(updated))
}

func (cc *CostController) DeleteCost(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	if err := cc.service.Delete(id, userID); err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cost deleted successfully"})
}
