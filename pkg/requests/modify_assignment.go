package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/pkg/terminal"
)

type modifyAssignmentInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type modifyAssignmentOutput struct {
	Message string `json:"message"`
}

func ModifyAssignment(appName, assignmentID, title, description string) error {
	inputs := modifyAssignmentInput{
		Title:       helpers.If[string](title != "", title, ""),
		Description: helpers.If[string](description != "", description, ""),
	}

	url := fmt.Sprintf("/submit/%s/%s", appName, assignmentID)

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	req, err := doRequest(http.MethodPut, url, &inputBody, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	if req.StatusCode == http.StatusBadRequest {
		terminal.PrintNotFoundResponse(appName)
		return err
	}

	op, err := decodeModifyAssignmentBody(req.Body)
	if err != nil {
		return err
	}

	terminal.PrintSuccessMsgResponse(op.Message)

	return nil
}

func decodeModifyAssignmentBody(body io.ReadCloser) (*modifyAssignmentOutput, error) {
	defer body.Close()
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse modifyAssignmentOutput
	if err := json.Unmarshal(result, &parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
