package validator

import "github.com/go-playground/validator/v10"

type ApiValidator struct {
	Validator *validator.Validate
}

func (cv *ApiValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
