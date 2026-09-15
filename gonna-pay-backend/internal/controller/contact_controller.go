package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/validation"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/httputil"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/service"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/request"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/view/response"
)

type ContactController struct {
	service service.ContactService
}

func NewContactController(service service.ContactService) *ContactController {
	return &ContactController{service: service}
}

// @Summary      Criar contato
// @Description  Cria um novo contato vinculado ao usuário autenticado
// @Tags         contacts
// @Accept       json
// @Produce      json
// @Param        body  body      request.CreateContactRequest  true  "Dados do contato"
// @Success      201   {object}  response.ContactResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      409   {object}  rest_errors.RestErrors
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /contact [post]
func (cc *ContactController) CreateContact(c *gin.Context) {
	ownerID := c.GetString("userID")

	var req request.CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	contact := domains.NewContact(ownerID, req.Name, req.Email, req.Phone, req.Category)

	created, err := cc.service.Create(contact)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewContactResponse(created))
}

// @Summary      Listar contatos
// @Description  Retorna todos os contatos do usuário autenticado
// @Tags         contacts
// @Produce      json
// @Param        page   query     int  false  "Número da página (padrão: 1)"
// @Param        limit  query     int  false  "Itens por página (padrão: 20, máximo: 100)"
// @Success      200  {object}  response.PaginatedResponse[response.ContactResponse]
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /contacts [get]
func (cc *ContactController) FindAllContacts(c *gin.Context) {
	ownerID := c.GetString("userID")

	page := 1
	limit := 20

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 100 {
			limit = parsed
		}
	}

	contacts, total, err := cc.service.FindAll(ownerID, page, limit)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewPaginatedResponse(
		response.NewContactResponseList(contacts),
		page, limit, total,
	))
}

// @Summary      Buscar contato por ID
// @Description  Retorna um contato pelo ID. Apenas o dono pode acessar
// @Tags         contacts
// @Produce      json
// @Param        id   path      string  true  "ID do contato"
// @Success      200  {object}  response.ContactResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors  "Contato não pertence ao usuário"
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /contact/{id} [get]
func (cc *ContactController) FindContactByID(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	contact, err := cc.service.FindByID(id, ownerID)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewContactResponse(contact))
}

// @Summary      Atualizar contato
// @Description  Atualiza os dados de um contato. Apenas o dono pode editar. Campos são opcionais (patch)
// @Tags         contacts
// @Accept       json
// @Produce      json
// @Param        id    path      string                        true  "ID do contato"
// @Param        body  body      request.UpdateContactRequest  true  "Dados para atualização"
// @Success      200   {object}  response.ContactResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      403   {object}  rest_errors.RestErrors
// @Failure      404   {object}  rest_errors.RestErrors
// @Failure      409   {object}  rest_errors.RestErrors
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /contact/{id} [patch]
func (cc *ContactController) UpdateContact(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	var req request.UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	contact := domains.NewContactWithID(id, ownerID, req.Name, req.Email, req.Phone, req.Category, time.Time{})

	updated, err := cc.service.Update(contact)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewContactResponse(updated))
}

// @Summary      Deletar contato
// @Description  Remove um contato. Apenas o dono pode deletar
// @Tags         contacts
// @Produce      json
// @Param        id   path      string  true  "ID do contato"
// @Success      200  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /contact/{id} [delete]
func (cc *ContactController) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	if err := cc.service.Delete(id, ownerID); err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "contact deleted successfully"})
}

// @Summary      Contatos por frequência
// @Description  Retorna os contatos que mais aparecem nos custos do usuário autenticado
// @Tags         contacts
// @Produce      json
// @Param        limit  query     int  false  "Número máximo de contatos (padrão: 5)"
// @Success      200    {array}   response.ContactFrequencyResponse
// @Failure      401    {object}  rest_errors.RestErrors
// @Failure      500    {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /contacts/frequency [get]
func (cc *ContactController) FindContactsByFrequency(c *gin.Context) {
	userID := c.GetString("userID")

	limit := 5
	if raw := c.Query("limit"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil {
			limit = parsed
		}
	}

	contacts, err := cc.service.FindContactsByFrequency(userID, limit)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewContactFrequencyResponseList(contacts))
}
