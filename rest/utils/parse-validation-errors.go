package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type ValidationErrors map[string]interface{}

func msgForField(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return fmt.Sprintf("Must be at max %s character(s) long", err.Param())
	case "email":
		return fmt.Sprint("Must be a valid email")
	case "uuid4":
		return fmt.Sprint("Must be a valid uuid")
	case "numeric":
		return fmt.Sprint("Must contain only numbers")
	case "oneof":
		return fmt.Sprintf("Value must be one of %s", err.Param())
	default:
		return err.Error()
	}
}

func ParseValidationErrors(err error) ValidationErrors {
	obj := ValidationErrors{}
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			obj[strcase.ToLowerCamel(fe.Field())] = msgForField(fe)
		}
	}

	return obj
}
