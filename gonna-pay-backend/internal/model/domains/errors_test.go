package domains_test

import (
	"errors"
	"testing"

	"github.com/rafabcanedo/gonna-pay/gonna-pay-backend/internal/model/domains"
	"github.com/stretchr/testify/assert"
)

func TestErrorSentinels(t *testing.T) {
	cases := []struct {
		name     string
		err      error
		sentinel error
	}{
		{"NewNotFoundError wraps ErrNotFound", domains.NewNotFoundError("msg"), domains.ErrNotFound},
		{"NewConflictError wraps ErrConflict", domains.NewConflictError("msg"), domains.ErrConflict},
		{"NewForbiddenError wraps ErrForbidden", domains.NewForbiddenError("msg"), domains.ErrForbidden},
		{"NewUnauthorizedError wraps ErrUnauthorized", domains.NewUnauthorizedError("msg"), domains.ErrUnauthorized},
		{"NewInvalidInputError wraps ErrInvalidInput", domains.NewInvalidInputError("msg"), domains.ErrInvalidInput},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ErrorIs(t, tc.err, tc.sentinel)
		})
	}
}

func TestErrorMessage(t *testing.T) {
	t.Run("Error() returns the custom message not the sentinel", func(t *testing.T) {
		err := domains.NewNotFoundError("user not found")

		assert.Equal(t, "user not found", err.Error())
	})
}

func TestErrorCrossContamination(t *testing.T) {
	t.Run("errors.Is does not match wrong sentinel", func(t *testing.T) {
		err := domains.NewNotFoundError("not found")

		assert.False(t, errors.Is(err, domains.ErrConflict))
		assert.False(t, errors.Is(err, domains.ErrForbidden))
	})
}
