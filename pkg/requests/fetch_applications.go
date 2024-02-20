package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/terminal"
)

type fetchApplicationsOutput struct {
	Applications []application.Application `json:"applications"`
}

func FetchApplications() error {

	res, err := doRequest(http.MethodGet, "/applications", nil, true)
	if err != nil {
		return err
	}

	op, err := decodeFetchApplications(res.Body)
	if err != nil {
		return err
	}

	runFetchApplicationsView(op.Applications)

	return nil
}

func decodeFetchApplications(body io.ReadCloser) (*fetchApplicationsOutput, error) {
	defer body.Close()
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse []application.Application
	if err := json.Unmarshal(response, &parsedResponse); err != nil {
		return nil, err
	}

	return &fetchApplicationsOutput{
		Applications: parsedResponse,
	}, nil
}

func runFetchApplicationsView(apps []application.Application) {
	terminal.ClearTerminal()

	terminal.Header()
	terminal.Gap()
	terminal.Title("All Applications")
	terminal.Gap()

	for idx, a := range apps {

		terminal.Content(fmt.Sprintf("%d. %s %s: (%s)", idx+1, a.ID, a.Name, a.CreatedAt.Format("2006/01/02")))
	}

	terminal.Gap()
	terminal.Footer()
}
