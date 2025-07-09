// Package validator provides custom validation functionality using go-playground/validator.
package validator

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
