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

type GroupController struct {
	service service.GroupService
}

func NewGroupController(service service.GroupService) *GroupController {
	return &GroupController{service: service}
}

// @Summary      Criar grupo
// @Description  Cria um novo grupo. É possível já adicionar membros (contatos) pelo campo memberIds
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        body  body      request.CreateGroupRequest  true  "Dados do grupo"
// @Success      201   {object}  response.GroupResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      409   {object}  rest_errors.RestErrors
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /group [post]
func (gc *GroupController) CreateGroup(c *gin.Context) {
	ownerID := c.GetString("userID")

	var req request.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	group := domains.NewGroup(ownerID, req.Name, req.Category)

	created, err := gc.service.Create(group, req.MemberIDs)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewGroupResponse(created))
}

// @Summary      Listar grupos
// @Description  Retorna todos os grupos do usuário autenticado com paginação
// @Tags         groups
// @Produce      json
// @Param        page      query     int     false  "Número da página (padrão: 1)"
// @Param        limit     query     int     false  "Itens por página (padrão: 20, máximo: 100)"
// @Param        category  query     string  false  "Filtrar por categoria"  Enums(Dinner, Lunch, Entertainment, Travel, Others)
// @Param        search    query     string  false  "Buscar por nome"
// @Success      200  {object}  response.PaginatedGroupResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /groups [get]
func (gc *GroupController) FindAllGroups(c *gin.Context) {
	ownerID := c.GetString("userID")
	page, limit := httputil.ParsePagination(c)

	filters := domains.GroupFilters{
		Category: c.Query("category"),
		Search:   c.Query("search"),
	}

	groups, total, err := gc.service.FindAll(ownerID, page, limit, filters)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewPaginatedResponse(
		response.NewGroupResponseList(groups),
		page, limit, total,
	))
}

// @Summary      Buscar grupo por ID
// @Description  Retorna um grupo com a lista de membros. Apenas o dono pode acessar
// @Tags         groups
// @Produce      json
// @Param        id   path      string  true  "ID do grupo"
// @Success      200  {object}  response.GroupDetailResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors  "Grupo não pertence ao usuário"
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /group/{id} [get]
func (gc *GroupController) FindGroupByID(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	group, err := gc.service.FindByID(id, ownerID)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewGroupDetailResponse(group))
}

// @Summary      Atualizar grupo
// @Description  Atualiza nome e/ou categoria do grupo. Apenas o dono pode editar
// @Tags         groups
// @Accept       json
// @Produce      json
// @Param        id    path      string                      true  "ID do grupo"
// @Param        body  body      request.UpdateGroupRequest  true  "Dados para atualização"
// @Success      200   {object}  response.GroupResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      403   {object}  rest_errors.RestErrors
// @Failure      404   {object}  rest_errors.RestErrors
// @Failure      409   {object}  rest_errors.RestErrors
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /group/{id} [patch]
func (gc *GroupController) UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	var req request.UpdateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	group := domains.NewGroupForUpdate(id, ownerID, req.Name, req.Category)

	updated, err := gc.service.Update(group, ownerID)
	if err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewGroupResponse(updated))
}

// @Summary      Deletar grupo
// @Description  Remove um grupo. Apenas o dono pode deletar
// @Tags         groups
// @Produce      json
// @Param        id   path      string  true  "ID do grupo"
// @Success      200  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /group/{id} [delete]
func (gc *GroupController) DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	ownerID := c.GetString("userID")

	if err := gc.service.Delete(id, ownerID); err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "group deleted successfully"})
}

// @Summary      Adicionar membro
// @Description  Adiciona um contato como membro do grupo. Apenas o dono do grupo pode adicionar
// @Tags         groups
// @Produce      json
// @Param        id         path      string  true  "ID do grupo"
// @Param        contactId  path      string  true  "ID do contato"
// @Success      201  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      409  {object}  rest_errors.RestErrors  "Membro já está no grupo"
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /group/{id}/member/{contactId} [post]
func (gc *GroupController) AddMember(c *gin.Context) {
	groupID := c.Param("id")
	contactID := c.Param("contactId")
	ownerID := c.GetString("userID")

	if err := gc.service.AddMember(groupID, contactID, ownerID); err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "member added successfully"})
}

// @Summary      Remover membro
// @Description  Remove um contato do grupo. Apenas o dono do grupo pode remover
// @Tags         groups
// @Produce      json
// @Param        id         path      string  true  "ID do grupo"
// @Param        contactId  path      string  true  "ID do contato"
// @Success      200  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      403  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /group/{id}/member/{contactId} [delete]
func (gc *GroupController) RemoveMember(c *gin.Context) {
	groupID := c.Param("id")
	contactID := c.Param("contactId")
	ownerID := c.GetString("userID")

	if err := gc.service.RemoveMember(groupID, contactID, ownerID); err != nil {
		httputil.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member removed successfully"})
}
