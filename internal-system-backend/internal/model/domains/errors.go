package domains

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
	ErrForbidden    = errors.New("forbidden")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

type appError struct {
	sentinel error
	msg      string
}

func (e *appError) Error() string { return e.msg }
func (e *appError) Unwrap() error { return e.sentinel }

func NewNotFoundError(msg string) error     { return &appError{ErrNotFound, msg} }
func NewConflictError(msg string) error     { return &appError{ErrConflict, msg} }
func NewForbiddenError(msg string) error    { return &appError{ErrForbidden, msg} }
func NewUnauthorizedError(msg string) error { return &appError{ErrUnauthorized, msg} }
func NewInvalidInputError(msg string) error { return &appError{ErrInvalidInput, msg} }
