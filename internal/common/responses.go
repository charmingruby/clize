package common

import "fmt"

type CreatedResponse struct {
	Message string `json:"message"`
}

func NewCreatedResponse(identifier string) CreatedResponse {
	return CreatedResponse{
		Message: fmt.Sprintf("%s created successfully", identifier),
	}
}
