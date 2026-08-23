package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/rest_errors"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/model/domains"
)

func RespondError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domains.ErrNotFound):
		c.JSON(http.StatusNotFound, rest_errors.NewNotFoundError(err.Error()))
	case errors.Is(err, domains.ErrForbidden):
		c.JSON(http.StatusForbidden, rest_errors.NewForbiddenError(err.Error()))
	case errors.Is(err, domains.ErrConflict):
		c.JSON(http.StatusConflict, rest_errors.NewConflictError(err.Error()))
	case errors.Is(err, domains.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, rest_errors.NewUnauthorizedRequestError(err.Error()))
	case errors.Is(err, domains.ErrInvalidInput):
		c.JSON(http.StatusBadRequest, rest_errors.NewBadRequestError(err.Error()))
	default:
		c.JSON(http.StatusInternalServerError, rest_errors.NewInternalServerError("internal server error"))
	}
}
