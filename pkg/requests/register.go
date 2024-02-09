package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	cliui "github.com/charmingruby/clize/pkg/cli_ui"
)

type registerInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerOutput struct {
	Message string `json:"message"`
}

func Register(username, email, password string) error {
	inputs := registerInput{
		Username: username,
		Email:    email,
		Password: password,
	}

	var inputBody bytes.Buffer
	if err := json.NewEncoder(&inputBody).Encode(inputs); err != nil {
		return err
	}

	res, err := doRequest(http.MethodPost, "/sign-up", &inputBody, false)
	if err != nil {
		cliui.PrintServerError()
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return errors.New("unable to create")
	}

	op, err := decodeBody(res.Body)
	if err != nil {
		return err
	}

	print(op.Message)

	return nil
}

func decodeBody(body io.ReadCloser) (*registerOutput, error) {
	defer body.Close()
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse registerOutput
	if err := json.Unmarshal(response, &parsedResponse); err != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
