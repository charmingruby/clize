package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

type submitAssignmentOutput struct {
	Message string `json:"message"`
}

func SubmitAssignment(appName, assignmentID string) error {
	url := fmt.Sprintf("/submit/%s/%s", appName, assignmentID)

	req, err := doRequest(http.MethodPut, url, nil, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	op, err := decodeSubmitAssignmentBody(req.Body)
	if err != nil {
		return err
	}

	terminal.PrintSuccessMsgResponse(op.Message)

	return nil
}

func decodeSubmitAssignmentBody(body io.ReadCloser) (*submitAssignmentOutput, error) {
	defer body.Close()
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse *submitAssignmentOutput
	if err := json.Unmarshal(result, &parsedResponse); err != nil {
		return nil, err
	}

	return parsedResponse, nil
}
