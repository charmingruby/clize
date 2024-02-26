package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/pkg/terminal"
)

type modifyApplicationInput struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

func ModifyApplication(appName, name, context string) error {
	inputs := modifyApplicationInput{
		Name:    helpers.If[string](name != "", fixInputSpacing(name), ""),
		Context: helpers.If[string](context != "", fixInputSpacing(context), ""),
	}

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	req, err := doRequest(http.MethodPut, fmt.Sprintf("/applications/%s", appName), &inputBody, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	data := decodeBody(req.Body)

	if req.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", data.Message)
	}

	terminal.PrintSuccessMsgResponse(data.Message)

	return nil
}
