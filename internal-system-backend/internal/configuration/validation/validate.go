package validation

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
	"github.com/rafabcanedo/basic-internal-system/internal-system-backend/internal/configuration/rest_errors"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateError(err error) *rest_errors.RestErrors {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(err, &jsonErr) {
		return rest_errors.NewBadRequestError("invalid field type")
	} else if errors.As(err, &jsonValidationError) {
		causes := make([]rest_errors.Causes, 0, len(jsonValidationError))
		for _, e := range jsonValidationError {
			causes = append(causes, rest_errors.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}
		return rest_errors.NewBadRequestValidationError("Some fields are invalid", causes)
	}

	return rest_errors.NewBadRequestError("error trying to convert fields")
}
