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

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

// @Summary      Criar usuário
// @Description  Cria um novo usuário (registro). Não requer autenticação
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body  body      request.CreateUserRequest  true  "Dados do usuário"
// @Success      201   {object}  response.UserResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      409   {object}  rest_errors.RestErrors  "Email já cadastrado"
// @Failure      500   {object}  rest_errors.RestErrors
// @Router       /user [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	user := domains.NewUser(req.Name, req.Email, req.Password, req.Phone)

	created, err := uc.service.Create(user)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewUserResponse(created))
}

// @Summary      Listar usuários
// @Description  Retorna todos os usuários cadastrados
// @Tags         users
// @Produce      json
// @Success      200  {array}   response.UserResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /users [get]
func (uc *UserController) FindAllUsers(c *gin.Context) {
	users, err := uc.service.FindAll()
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponseList(users))
}

// @Summary      Buscar usuário por ID
// @Description  Retorna um usuário pelo ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "ID do usuário"
// @Success      200  {object}  response.UserResponse
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /user/{id} [get]
func (uc *UserController) FindUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := uc.service.FindByID(id)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(user))
}

// @Summary      Atualizar usuário
// @Description  Atualiza os dados de um usuário pelo ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      string                    true  "ID do usuário"
// @Param        body  body      request.UpdateUserRequest  true  "Dados para atualização"
// @Success      200   {object}  response.UserResponse
// @Failure      400   {object}  rest_errors.RestErrors
// @Failure      401   {object}  rest_errors.RestErrors
// @Failure      404   {object}  rest_errors.RestErrors
// @Failure      409   {object}  rest_errors.RestErrors  "Email já em uso"
// @Failure      500   {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /user/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var req request.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := validation.ValidateError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	user := domains.NewUserWithID(id, req.Name, req.Email, req.Password, req.Phone)

	updated, err := uc.service.Update(user)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(updated))
}

// @Summary      Deletar usuário
// @Description  Remove um usuário pelo ID
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "ID do usuário"
// @Success      200  {object}  map[string]string  "message"
// @Failure      401  {object}  rest_errors.RestErrors
// @Failure      404  {object}  rest_errors.RestErrors
// @Failure      500  {object}  rest_errors.RestErrors
// @Security     CookieAuth
// @Router       /user/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
