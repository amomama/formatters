package formatter

import (
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"net/http"
	"strings"
)

type Error struct {
	Attribute string `json:"attribute"`
	Validator `json:"validator"`
}

type Validator struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

type ValidationErrors struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

func (errors *ValidationErrors) AddError(err Error) {
	errors.Errors = append(errors.Errors, err)
}

func ValidationResponse(err error) (code int, errors ValidationErrors) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		code = http.StatusInternalServerError
		errors.Message = http.StatusText(code)
		return
	}

	code = http.StatusUnprocessableEntity
	errors.Message = http.StatusText(code)
	for _, err := range err.(validator.ValidationErrors) {
		errors.AddError(Error{
			Attribute: strings.ToLower(err.Field()),
			Validator: Validator{
				Name:  strcase.ToLowerCamel(err.Tag()),
				Value: err.Param(),
			},
		})
	}

	return
}
