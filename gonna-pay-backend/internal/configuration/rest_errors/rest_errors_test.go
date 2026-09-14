package rest_errors_test

import (
	"net/http"
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/rest_errors"
	"github.com/stretchr/testify/assert"
)

func TestRestErrorConstructors(t *testing.T) {
	cases := []struct {
		name    string
		result  *rest_errors.RestErrors
		code    int
		err     string
		message string
	}{
		{
			name:    "NewBadRequestError returns code 400 and err bad_request",
			result:  rest_errors.NewBadRequestError("bad request"),
			code:    http.StatusBadRequest,
			err:     "bad_request",
			message: "bad request",
		},
		{
			name:    "NewInternalServerError returns code 500 and err internal_server_error",
			result:  rest_errors.NewInternalServerError("internal error"),
			code:    http.StatusInternalServerError,
			err:     "internal_server_error",
			message: "internal error",
		},
		{
			name:    "NewNotFoundError returns code 404 and err not_found",
			result:  rest_errors.NewNotFoundError("not found"),
			code:    http.StatusNotFound,
			err:     "not_found",
			message: "not found",
		},
		{
			name:    "NewConflictError returns code 409 and err conflict",
			result:  rest_errors.NewConflictError("conflict"),
			code:    http.StatusConflict,
			err:     "conflict",
			message: "conflict",
		},
		{
			name:    "NewForbiddenError returns code 403 and err forbidden",
			result:  rest_errors.NewForbiddenError("forbidden"),
			code:    http.StatusForbidden,
			err:     "forbidden",
			message: "forbidden",
		},
		{
			name:    "NewUnauthorizedRequestError returns code 401 and err unauthorized",
			result:  rest_errors.NewUnauthorizedRequestError("unauthorized"),
			code:    http.StatusUnauthorized,
			err:     "unauthorized",
			message: "unauthorized",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.code, tc.result.Code)
			assert.Equal(t, tc.err, tc.result.Err)
			assert.Equal(t, tc.message, tc.result.Message)
		})
	}
}

func TestNewBadRequestValidationError(t *testing.T) {
	t.Run("includes causes", func(t *testing.T) {
		causes := []rest_errors.Causes{
			{Field: "name", Message: "name is required"},
		}

		result := rest_errors.NewBadRequestValidationError("Some fields are invalid", causes)

		assert.Equal(t, http.StatusBadRequest, result.Code)
		assert.Equal(t, "bad_request", result.Err)
		assert.Equal(t, "Some fields are invalid", result.Message)
		assert.Len(t, result.Causes, 1)
		assert.Equal(t, "name", result.Causes[0].Field)
		assert.Equal(t, "name is required", result.Causes[0].Message)
	})
}

func TestRestErrors_Error(t *testing.T) {
	t.Run("Error() returns the message", func(t *testing.T) {
		err := rest_errors.NewNotFoundError("user not found")

		assert.Equal(t, "user not found", err.Error())
	})
}
