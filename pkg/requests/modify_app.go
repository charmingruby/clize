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

type modifyApplicationInput struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

type modifyApplicationOutput struct {
	Message string `json:"message"`
}

func ModifyApp(appName, name, context string) error {
	inputs := modifyApplicationInput{
		Name:    helpers.If[string](name != "", name, ""),
		Context: helpers.If[string](context != "", context, ""),
	}

	println("aqui1")

	println(appName)
	println(name)
	println(context)

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	println("aqui2")

	req, err := doRequest(http.MethodPut, fmt.Sprintf("/applications/%s", appName), &inputBody, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	println("aqui3")

	if req.StatusCode == http.StatusBadRequest {
		terminal.PrintNotFoundResponse(appName)
		return err
	}

	println("aqui4")

	op, err := decodeModifyAppBody(req.Body)
	if err != nil {
		return err
	}

	println("aqui5")

	terminal.PrintSuccessMsgResponse(op.Message)

	return nil
}

func decodeModifyAppBody(body io.ReadCloser) (*modifyApplicationOutput, error) {
	defer body.Close()
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse modifyApplicationOutput
	if err := json.Unmarshal(result, &parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
