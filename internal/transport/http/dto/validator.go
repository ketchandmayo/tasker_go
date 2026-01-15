package dto

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationError(err error) string {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, fe := range ve {
			if fe.Tag() == "required" {
				return fe.Field() + " is required"
			} else if fe.Tag() == "max" {
				return fe.Field() + " must be at most 100 characters"
			} else {
				return fe.Field() + " invalid validation"
			}
		}
	}

	return "invalid request"
}
