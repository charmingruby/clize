package endpoints

import "fmt"

type Response[d any] struct {
	Data    *d     `json:"data,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WrapResponse[d any](data *d, status int, message string) *Response[d] {
	res := &Response[d]{
		Data:    data,
		Status:  status,
		Message: message,
	}

	return res
}

func NewAddItemResponse(item, aggregate string) string {
	return fmt.Sprintf("%s added successfully to %s", item, aggregate)
}

func NewRemovedItemResponse(item, aggregate string) string {
	return fmt.Sprintf("%s removed successfully from %s", item, aggregate)
}

func NewModifiedResponse(identifier string) string {
	return fmt.Sprintf("%s modified successfully", identifier)
}

func NewCreatedResponse(identifier string) string {
	return fmt.Sprintf("%s created successfully", identifier)
}

func NewDeletedResponse(identifier string) string {
	return fmt.Sprintf("%s deleted successfully", identifier)
}

func NewFetchedResponse(entity string, quantity int) string {
	return fmt.Sprintf("%d %ss found", quantity, entity)
}

func NewCredentialsDoesntMatchResponse() string {
	return "Credentials does not match"
}
