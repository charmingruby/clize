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

func ModifyAssignment(appName, assignmentTitle, title, description string) error {
	inputs := modifyAssignmentInput{
		Title:       helpers.If[string](title != "", title, ""),
		Description: helpers.If[string](description != "", description, ""),
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

	statusCode := res.StatusCode
	if statusCode != http.StatusOK {
		if statusCode != http.StatusNotFound {
			errRes := decodeNotFoundError(res.Body)
			terminal.PrintErrorResponse(errRes.Message)
			return err
		}

		badRequestMsg := fmt.Sprintf("Error modifiying %s in %s", assignmentTitle, appName)
		terminal.PrintErrorResponse(badRequestMsg)
		return err
	}

	op, err := decodeModifyAssignmentBody(res.Body)
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
