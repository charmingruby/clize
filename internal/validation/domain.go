package validation

import (
	"fmt"

	"github.com/charmingruby/clize/helpers"
)

type GenericValidationError struct {
	Message string
}

func (e *GenericValidationError) Error() string {
	return e.Message
}

type RequiredFieldError struct {
	FieldName string
}

func (e *RequiredFieldError) Error() string {
	msg := fmt.Sprintf("%s is required", e.FieldName)
	return msg
}

type FieldLengthError struct {
	FieldName      string
	IsMinimumError bool
	Quantity       int
}

func (e *FieldLengthError) Error() string {
	minMsg := fmt.Sprintf("%s must be at least %d characters", e.FieldName, e.Quantity)
	maxMsg := fmt.Sprintf("%s must be a maximum of %d characters", e.FieldName, e.Quantity)

	msg := helpers.If[string](
		e.IsMinimumError,
		minMsg,
		maxMsg,
	)

	return msg
}
