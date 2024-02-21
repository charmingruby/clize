package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/charmingruby/clize/pkg/terminal"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token"`
}

func Auth(username, password, path string) error {
	creds := credentials{username, password}

	var credsBody bytes.Buffer
	if err := json.NewEncoder(&credsBody).Encode(creds); err != nil {
		return err
	}

	res, err := doRequest("POST", path, &credsBody, false)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	return createTokenCache(res.Body)
}

type cacheToken struct {
	Token string `json:"token"`
}

func createTokenCache(body io.ReadCloser) error {
	defer body.Close()
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	var result authResponse
	if err := json.Unmarshal(response, &result); err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(".cacheToken")
	if err != nil {

		return err
	}

	cache := &cacheToken{
		Token: result.Token,
	}
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	msg := "Authenticated successfully"
	terminal.PrintSuccessMsgResponse(msg)

	return nil
}

func readCacheToken() (string, error) {
	data, err := os.ReadFile(".cacheToken")
	if err != nil {
		return "", err
	}

	var cache cacheToken
	err = json.Unmarshal(data, &cache)
	if err != nil {
		return "", err
	}

	return cache.Token, nil
}
