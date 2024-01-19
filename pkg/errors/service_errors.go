package errors

import "fmt"

type UniqueValueViolationError struct {
	Entity  string `json:"entity"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewUniqueValueViolationErrorMessage(entity string, field string) string {
	return fmt.Sprintf("%s %s must be unique", entity, field)
}

func (v *UniqueValueViolationError) Error() string {
	return NewUniqueValueViolationErrorMessage(v.Entity, v.Field)
}
