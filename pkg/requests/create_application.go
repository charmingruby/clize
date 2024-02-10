package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/pkg/terminal"
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
	terminal.Header()
	terminal.Gap()

	terminal.Padding()
	terminal.BoldGreen.Printf(op.Message)

	terminal.Gap()
	terminal.Footer()
}
