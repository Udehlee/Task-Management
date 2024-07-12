package utils

import "github.com/go-playground/validator/v10"

func NewValidate() *validator.Validate {
	return validator.New()
}

func Validate(s interface{}) error {
	validate := NewValidate()
	return validate.Struct(s)
}
