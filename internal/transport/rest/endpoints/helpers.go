package endpoints

import (
	"fmt"

	"github.com/charmingruby/clize/helpers"
)

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
	return fmt.Sprintf("\"%s\" added successfully to \"%s\"", item, aggregate)
}

func NewRemovedItemResponse(item, aggregate string) string {
	return fmt.Sprintf("\"%s\" removed successfully from \"%s\"", item, aggregate)
}

func NewModifiedResponse(identifier string) string {
	return fmt.Sprintf("\"%s\" modified successfully", identifier)
}

func NewCreatedResponse(identifier string) string {
	return fmt.Sprintf("\"%s\" created successfully", identifier)
}

func NewDeletedResponse(identifier string) string {
	return fmt.Sprintf("\"%s\" deleted successfully", identifier)
}

func NewFetchedResponse(entity string, quantity int) string {
	single := quantity == 1 || quantity == 0
	msg := helpers.If[string](single, fmt.Sprintf("%d \"%s\" found", quantity, entity), fmt.Sprintf("%d %ss found", quantity, entity))

	return msg
}

func NewCredentialsDoesntMatchResponse() string {
	return "Credentials does not match"
}
