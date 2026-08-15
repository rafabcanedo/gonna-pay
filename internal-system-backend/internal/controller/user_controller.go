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

func (uc *UserController) FindAllUsers(c *gin.Context) {
	users, err := uc.service.FindAll()
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponseList(users))
}

func (uc *UserController) FindUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := uc.service.FindByID(id)
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.NewUserResponse(user))
}

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

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := uc.service.Delete(id); err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
