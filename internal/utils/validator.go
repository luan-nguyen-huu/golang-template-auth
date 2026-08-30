package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct validates a struct based on its `validate` tags.
func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		if valErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsgs []string
			for _, e := range valErrors {
				errMsgs = append(errMsgs, fmt.Sprintf("field '%s' failed validation '%s'", e.Field(), e.Tag()))
			}
			return fmt.Errorf("%s", strings.Join(errMsgs, ", "))
		}
		return err
	}
	return nil
}
