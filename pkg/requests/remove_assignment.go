package requests

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

type removeAssignmentInput struct {
	AppName    string
	Assignment string
}

func RemoveAssignment(appName, assignmentName string) error {
	inputs := removeAssignmentInput{appName, assignmentName}

	url := fmt.Sprintf("/assignments/%s/%s", inputs.AppName, inputs.Assignment)

	res, err := doRequest(http.MethodDelete, url, nil, true)
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
