package requests

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

func Ping() error {
	_, err := doRequest(http.MethodGet, "/health-check", nil, false)

	if err != nil {
		serverDownErr := fmt.Errorf("server is not running")
		return serverDownErr
	}

	terminal.PrintSuccessMsgResponse("server is running")
	return nil
}
