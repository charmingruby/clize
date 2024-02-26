package requests

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

func SubmitAssignment(appName, assignmentID string) error {
	url := fmt.Sprintf("/submit/%s/%s", appName, assignmentID)

	res, err := doRequest(http.MethodPut, url, nil, true)
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
