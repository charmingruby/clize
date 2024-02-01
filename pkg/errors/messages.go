package errors

import "fmt"

func NewUniqueValueViolationErrorMessage(entity string, field string) string {
	return fmt.Sprintf("%s %s must be unique", entity, field)
}

func NewResourceNotFoundErrorMessage(entity string) string {
	msg := fmt.Sprintf("%s not found", entity)

	return msg
}

func NewNotNullableErrorMessage(fieldsOption []string) string {
	var msg string
	var fieldsStr string

	for idx, field := range fieldsOption {

		if idx+1 == len(fieldsOption) {

			fieldsStr += fmt.Sprintf(" or %s", field)

		} else {
			fieldsStr += fmt.Sprintf(", %s", field)
		}

	}

	msg = fmt.Sprintf("Body cannot be empty, please provide at least: %s.", fieldsStr)

	return msg
}
