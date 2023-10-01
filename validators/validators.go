package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
