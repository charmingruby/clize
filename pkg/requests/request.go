package requests

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func doRequest(
	method, path string, body io.Reader, auth bool,
) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost:8080%s", path)

	req, err := http.NewRequest(method, url, body)
	if err != nil {

		return nil, err
	}

	if auth {
		token, err := readCacheToken()

		if err != nil {
			log.Println("Cannot read cache token.")
			return nil, err
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	return http.DefaultClient.Do(req)
}
