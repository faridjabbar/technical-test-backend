package helper

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterValidation(validate *validator.Validate) {
	validate.RegisterValidation("role", validateRole)
}

func validateRole(fieldLevel validator.FieldLevel) bool {
	value := fieldLevel.Field().String()

	const pattern = "(Administrator|Client)"
	regex := regexp.MustCompile(pattern)
	result := regex.MatchString(value)
	return result
}
