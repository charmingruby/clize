package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

type addAssignmentInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type addAssignmentOutput struct {
	Message string `json:"message"`
}

func AddAssignment(appName, title, description string) error {
	inputs := addAssignmentInput{title, description}

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	req, err := doRequest(http.MethodPost, fmt.Sprintf("/applications/%s/assignments", appName), &inputBody, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	if req.StatusCode == http.StatusBadRequest {
		terminal.PrintNotFoundResponse(appName)
		return err
	}

	op, err := decodeAddAssignment(req.Body)
	if err != nil {
		return err
	}

	terminal.PrintSuccessMsgResponse(op.Message)

	return nil
}

func decodeAddAssignment(body io.ReadCloser) (*addAssignmentOutput, error) {
	defer body.Close()
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse addAssignmentOutput
	if err := json.Unmarshal(result, &parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
