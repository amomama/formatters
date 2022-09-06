package formatters

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"net/http"
	"testing"
)

func TestValidationResponse(t *testing.T) {
	request := struct {
		Name     string `json:"name" validate:"required,min=4"`
		Email    string `json:"email" validate:"required,email,min=4,max=100"`
		Password string `json:"password" validate:"required"`
	}{
		Name:     "te",
		Email:    "test",
		Password: "",
	}

	validation := validator.New()
	validationError := validation.Struct(request)

	assert.NotEqual(t, nil, validationError)

	code, response := ValidationResponse(validationError)
	jsonResponse, _ := json.Marshal(response)

	expected := ValidationErrors{
		Message: "Unprocessable Entity",
		Errors: []Error{
			{
				Attribute: "name",
				ValidationAttributes: ValidationAttributes{
					Name:  "min",
					Value: "4",
				},
			},
			{
				Attribute: "email",
				ValidationAttributes: ValidationAttributes{
					Name: "email",
				},
			},
			{
				Attribute: "password",
				ValidationAttributes: ValidationAttributes{
					Name: "required",
				},
			},
		},
	}
	jsonExpected, _ := json.Marshal(expected)

	assert.Equal(t, code, http.StatusUnprocessableEntity)
	assert.Equal(t, string(jsonResponse), string(jsonExpected))
}
