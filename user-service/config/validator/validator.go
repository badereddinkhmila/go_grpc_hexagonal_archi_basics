package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	valid *validator.Validate = initValidator()
)

func initValidator() *validator.Validate {
	return validator.New()
}

func GetValidator() *validator.Validate {
	return valid
}

func ValidateStruct(obj any) error {
	validator := GetValidator()
	err := validator.Struct(obj)
	return err
}
