package validation_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/configuration/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateError(t *testing.T) {
	t.Run("json unmarshal type error returns bad request", func(t *testing.T) {
		var target struct{ Name int }
		err := json.Unmarshal([]byte(`{"name":"notanint"}`), &target)

		result := validation.ValidateError(err)

		assert.Equal(t, http.StatusBadRequest, result.Code)
		assert.Equal(t, "invalid field type", result.Message)
	})

	t.Run("validation errors returns bad request with causes", func(t *testing.T) {
		type req struct {
			Name string `validate:"required"`
		}
		err := validator.New().Struct(req{})

		result := validation.ValidateError(err)

		assert.Equal(t, http.StatusBadRequest, result.Code)
		assert.Equal(t, "Some fields are invalid", result.Message)
		assert.NotEmpty(t, result.Causes)
		assert.Equal(t, "Name", result.Causes[0].Field)
	})

	t.Run("unknown error returns generic bad request", func(t *testing.T) {
		result := validation.ValidateError(errors.New("something unexpected"))

		assert.Equal(t, http.StatusBadRequest, result.Code)
		assert.Equal(t, "error trying to convert fields", result.Message)
	})
}
