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

type InvalidPayloadError struct {
	Message        string   `json:"message"`
	RequiredFields []string `json:"required_fields"`
}

func (ip *InvalidPayloadError) Error() string {
	return ip.Message
}

func NewInvalidPayloadErrorMessage(requiredFields []string) string {
	var msg string
	var fieldsStr string

	for idx, field := range requiredFields {

		if idx+1 == len(requiredFields) {

			fieldsStr += fmt.Sprintf(" and %s", field)

		} else if idx == 0 {
			fieldsStr += fmt.Sprint(field)
		} else {
			fieldsStr += fmt.Sprintf(", %s", field)
		}
	}

	msg = fmt.Sprintf("Invalid payload, fields required: %s.", fieldsStr)

	return msg
}
