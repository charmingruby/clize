package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

type removeAssignmentInput struct {
	AppName    string
	Assignment string
}

type removeAssignmentOutput struct {
	Message string `json:"message"`
}

func RemoveAssignment(appName, assignmentName string) error {
	inputs := removeAssignmentInput{appName, assignmentName}

	url := fmt.Sprintf("/assignments/%s/%s", inputs.AppName, inputs.Assignment)

	res, err := doRequest(http.MethodDelete, url, nil, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	statusCode := res.StatusCode
	if statusCode != http.StatusOK {
		if statusCode == http.StatusNotFound {
			errRes := decodeNotFoundError(res.Body)
			terminal.PrintErrorResponse(errRes)
			return err
		}

		if statusCode == http.StatusBadRequest {
			return err
		}
	}

	op, err := decodeRemoveAssignmentBody(res.Body)
	if err != nil {
		return err
	}

	terminal.PrintSuccessMsgResponse(op.Message)

	return nil
}

func decodeRemoveAssignmentBody(body io.ReadCloser) (*removeAssignmentOutput, error) {
	defer body.Close()
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse removeAssignmentOutput
	if err := json.Unmarshal(result, &parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
