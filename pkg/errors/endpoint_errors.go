package errors

import "fmt"

type NotNullableBodyError struct {
	Message string   `json:"message"`
	Fields  []string `json:"fields"`
}

func (e *NotNullableBodyError) Error() string {
	return NewNotNullableErrorMessage(e.Fields)
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
