package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/pkg/terminal"
)

type modifyAssignmentInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func ModifyAssignment(appName, assignmentTitle, title, description string) error {
	inputs := modifyAssignmentInput{
		Title:       helpers.If[string](title != "", fixInputSpacing(title), ""),
		Description: helpers.If[string](description != "", fixInputSpacing(description), ""),
	}

	url := fmt.Sprintf("/assignments/%s/%s", appName, assignmentTitle)

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	res, err := doRequest(http.MethodPut, url, &inputBody, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	data := decodeBody(res.Body)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", data.Message)
	}

	terminal.PrintSuccessMsgResponse(data.Message)

	return nil
}
