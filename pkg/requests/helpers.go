package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

type uncheckedResponse struct {
	Data    any    `json:"data,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func decodeBody(body io.ReadCloser) *uncheckedResponse {
	defer body.Close()
	result, _ := ioutil.ReadAll(body)

	var response uncheckedResponse
	json.Unmarshal(result, &response)

	return &response
}

type genericResponse[v any] struct {
	Data    v      `json:"data,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func decodeBodyWithInterface[v any](body io.ReadCloser) *genericResponse[v] {
	defer body.Close()
	result, _ := ioutil.ReadAll(body)

	var response genericResponse[v]
	json.Unmarshal(result, &response)

	return &response
}

func fixInputSpacing(input string) string {
	input = strings.ReplaceAll(input, "_", " ")
	return input
}
