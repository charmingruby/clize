package requests

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/terminal"
)

func FetchApplications() error {
	res, err := doRequest(http.MethodGet, "/applications", nil, true)
	if err != nil {
		return err
	}

	data := decodeBodyWithInterface[[]application.Application](res.Body)

	runFetchApplicationsView(data.Data)

	return nil
}

func runFetchApplicationsView(apps []application.Application) {

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
