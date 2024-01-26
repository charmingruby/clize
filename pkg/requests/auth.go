package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func Auth(path string) error {
	res, err := doRequest("GET", path, nil, false)
	if err != nil {
		return err
	}

	return createTokenCache(res.Body)
}

type cacheToken struct {
	Token string `json:"token"`
}

func createTokenCache(body io.ReadCloser) error {
	token, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	file, err := os.Create(".cacheToken")
	if err != nil {
		return err
	}

	cache := cacheToken{string(token)}
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

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
