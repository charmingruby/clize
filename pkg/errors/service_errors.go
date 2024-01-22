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
	return v.Message
}

type ResourceNotFoundError struct {
	Entity  string `json:"entity"`
	Message string `json:"message"`
}

func (nf *ResourceNotFoundError) Error() string {
	return nf.Message
}

func NewResourceNotFoundErrorMessage(entity string) string {
	msg := fmt.Sprintf("%s not found", entity)

	return msg
}
