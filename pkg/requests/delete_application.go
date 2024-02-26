package requests

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

func DeleteApplication(name string) error {
	res, err := doRequest(http.MethodDelete, fmt.Sprintf("/applications/%s", name), nil, true)
	if err != nil {
		return err
	}

	data := decodeBody(res.Body)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", data.Message)
	}

	terminal.PrintSuccessMsgResponse(data.Message)

	return nil
}
