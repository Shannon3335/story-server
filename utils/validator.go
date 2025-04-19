package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct(s any) error {
	err := validate.Struct(s)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			return validationErrs
		}
		return err
	}
	return nil
}
