package requests

import (
	"encoding/json"
	"fmt"
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

func makeExcerpt(content string) string {
	chunkSize := 20

	var chunks []string

	if len(content) < chunkSize {
		return content
	}

	for i := 0; i < chunkSize; i += chunkSize {
		end := i + chunkSize

		if i == 17 {
			break
		}

		chunks = append(chunks, content[i:end])
	}

	formattedResult := strings.Join(chunks, "")

	return fmt.Sprintf("%s...", formattedResult)
}
