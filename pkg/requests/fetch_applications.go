package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/internal/app/domain/application"
	cliui "github.com/charmingruby/clize/pkg/cli_ui"
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
	cliui.Header()
	cliui.Gap()
	cliui.Title("All Applications")
	cliui.Gap()

	for idx, a := range apps {

		cliui.Content(fmt.Sprintf("%d. %s %s: (%s)", idx+1, a.ID, a.Name, a.CreatedAt.Format("2006/01/02")))
	}

	cliui.Gap()
	cliui.Footer()
}
