package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
)

type createApplicationInput struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

func CreateApplication(name, context string) error {
	inputs := createApplicationInput{
		Name:    name,
		Context: context,
	}

	inputs.Name = fixInputSpacing(inputs.Name)
	inputs.Context = fixInputSpacing(inputs.Context)

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	res, err := doRequest(http.MethodPost, "/applications", &inputBody, true)
	if err != nil {
		return err
	}

	data := decodeBody(res.Body)

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("%s", data.Message)
	}

	terminal.PrintSuccessMsgResponse(data.Message)

	return nil
}
