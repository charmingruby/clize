package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/fatih/color"
)

type createApplicationInput struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

type createApplicationOutput struct {
	Message string `json:"message"`
}

func CreateApplication(name, context string) error {
	inputs := createApplicationInput{
		Name:    name,
		Context: context,
	}

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	res, err := doRequest(http.MethodPost, "/applications", &inputBody, true)
	if err != nil {
		return err
	}

	op, err := decodeCreateApplicationBody(res.Body)
	if err != nil {
		return err
	}

	runCreateApplicationView(op)

	return nil
}

func decodeCreateApplicationBody(body io.ReadCloser) (*createApplicationOutput, error) {
	defer body.Close()
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse createApplicationOutput
	if err := json.Unmarshal(response, &parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}

func runCreateApplicationView(op *createApplicationOutput) {
	cliui.Header()
	cliui.Gap()

	cliui.Padding()
	color.Green(op.Message)

	cliui.Gap()
	cliui.Footer()
}