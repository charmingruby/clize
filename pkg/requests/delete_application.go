package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

type deleteApplicationOutput struct {
	Message string `json:"message"`
}

func DeleteApplication(name string) error {
	res, err := doRequest(http.MethodDelete, fmt.Sprintf("/applications/%s", name), nil, true)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusNotFound {
		terminal.PrintNotFoundResponse(name)
		return err
	}

	op, err := decodeDeleteApplicationBody(res.Body)
	if err != nil {
		return err
	}

	terminal.PrintSuccessMsgResponse(op.Message)

	return nil
}

func decodeDeleteApplicationBody(body io.ReadCloser) (*deleteApplicationOutput, error) {
	defer body.Close()
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse deleteApplicationOutput
	if err := json.Unmarshal(result, &parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
