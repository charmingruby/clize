package requests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/charmingruby/clize/internal/validation"
)

func decodeNotFoundError(body io.ReadCloser) *validation.ResourceNotFoundError {
	defer body.Close()
	result, _ := ioutil.ReadAll(body)

	var parsedError validation.ResourceNotFoundError
	json.Unmarshal(result, &parsedError)

	return &parsedError
}

func putCorrectSpacingOnInputs(input string) string {
	input = strings.ReplaceAll(input, "_", " ")
	return input
}
