package utils

import (
	"github.com/go-playground/validator/v10"
)

type ValidatorError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

var Validator = validator.New()

func ParserValidation(input interface{}) []*ValidatorError {
	var fieldErrors []*ValidatorError

	err := Validator.Struct(input)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidatorError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()

			fieldErrors = append(fieldErrors, &el)
		}

		return fieldErrors
	}

	return nil
}
