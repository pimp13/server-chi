package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func validateStruct(data interface{}) []string {
	err := Validate.Struct(data)
	if err == nil {
		return nil
	}

	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		fmt.Println(err)
		return nil
	}

	var validationErrors []string
	var validateErrs validator.ValidationErrors
	if errors.As(err, &validateErrs) {
		for _, e := range validateErrs {
			validationErrors = append(validationErrors, fmt.Sprintf(
				"فیلد %s نامعتبر است: قانون %s=%v نقض شده است",
				e.Field(), e.Tag(), e.Param(),
			))
		}
	}

	return validationErrors
}
