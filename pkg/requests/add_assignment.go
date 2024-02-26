package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

type addAssignmentInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func AddAssignment(appName, title, description string) error {
	inputs := addAssignmentInput{
		Title:       title,
		Description: description,
	}

	inputs.Description = fixInputSpacing(inputs.Description)

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	req, err := doRequest(http.MethodPost, fmt.Sprintf("/applications/%s/assignments", appName), &inputBody, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	data := decodeBody(req.Body)

	if req.StatusCode != http.StatusCreated {
		return fmt.Errorf("%s", data.Message)
	}

	terminal.PrintSuccessMsgResponse(data.Message)

	return nil
}
