package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/validation"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/httputil"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/service"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/request"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/response"
)

type CostController struct {
	service service.CostService
}

func NewCostController(service service.CostService) *CostController {
	return &CostController{service: service}
}

// @Summary      Criar custo
// @Description  Cria um novo custo. Se groupId for informado, o valor é dividido entre os membros do grupo. ownerPercentage é opcional — se omitido, a divisão é igualitária
// @Tags         costs
// @Accept       json
// @Produce      json
// @Param        body  body      request.CreateCostRequest  true  "Dados do custo"
// @Success      201   {object}  response.CostDetailResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      404   {object}  rest_errors.RestErrors  "Grupo não encontrado"
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /cost [post]
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
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewCostDetailResponse(created))
}

// @Summary      Listar custos
// @Description  Retorna todos os custos do usuário autenticado (sem detalhes dos splits)
// @Tags         costs
// @Produce      json
// @Success      200  {array}   response.CostResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /costs [get]
func (cc *CostController) FindAllCosts(c *gin.Context) {
	userID := c.GetString("userID")

	costs, err := cc.service.FindAll(userID)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewCostResponseList(costs))
}

// @Summary      Buscar custo por ID
// @Description  Retorna um custo com os detalhes de splits. Apenas o dono pode acessar
// @Tags         costs
// @Produce      json
// @Param        id   path      string  true  "ID do custo"
// @Success      200  {object}  response.CostDetailResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors  "Custo não pertence ao usuário"
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /cost/{id} [get]
func (cc *CostController) FindCostByID(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	cost, err := cc.service.FindByID(id, userID)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewCostDetailResponse(cost))
}

// @Summary      Atualizar custo
// @Description  Atualiza os dados de um custo. Apenas o dono pode editar. Os splits são recalculados automaticamente
// @Tags         costs
// @Accept       json
// @Produce      json
// @Param        id    path      string                    true  "ID do custo"
// @Param        body  body      request.UpdateCostRequest  true  "Dados para atualização"
// @Success      200   {object}  response.CostDetailResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      403   {object}  rest_errors.RestErrors
// @Failure      404   {object}  rest_errors.RestErrors
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /cost/{id} [patch]
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
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewCostDetailResponse(updated))
}

// @Summary      Deletar custo
// @Description  Remove um custo. Apenas o dono pode deletar
// @Tags         costs
// @Produce      json
// @Param        id   path      string  true  "ID do custo"
// @Success      200  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /cost/{id} [delete]
func (cc *CostController) DeleteCost(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userID")

	if err := cc.service.Delete(id, userID); err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "cost deleted successfully"})
}
