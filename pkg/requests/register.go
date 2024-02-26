package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	terminal "github.com/charmingruby/clize/pkg/terminal"
)

type registerInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
		terminal.PrintServerError()
		return err
	}

	data := decodeBody(res.Body)

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("%s", data.Message)
	}

	terminal.PrintSuccessMsgResponse(data.Message)

	return nil
}
