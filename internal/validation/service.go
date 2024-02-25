package validation

import "fmt"

type UniqueValueViolationError struct {
	Entity  string `json:"entity"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v *UniqueValueViolationError) Error() string {
	return v.Message
}

type ResourceNotFoundError struct {
	Entity string `json:"entity"`
}

func (nf *ResourceNotFoundError) Error() string {
	return NewResourceNotFoundErrorMessage(nf.Entity)
}

func NewUniqueValueViolationErrorMessage(entity string, field string) string {
	return fmt.Sprintf("%s %s must be unique", entity, field)
}

func NewResourceNotFoundErrorMessage(entity string) string {
	msg := fmt.Sprintf("%s not found", entity)

	return msg
}
