package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/charmingruby/clize/pkg/errors"
)

func decodeNotFoundError(body io.ReadCloser) *errors.ResourceNotFoundError {
	defer body.Close()
	result, _ := ioutil.ReadAll(body)

	var parsedError errors.ResourceNotFoundError
	json.Unmarshal(result, &parsedError)

	return &parsedError

}
