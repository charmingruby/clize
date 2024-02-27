package requests

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/terminal"
)

type getApplicationInput struct {
	Name string `json:"name"`
}

func GetApplication(name string) error {
	inputs := &getApplicationInput{
		Name: name,
	}

	res, err := doRequest(http.MethodGet, fmt.Sprintf("/applications/%s", inputs.Name), nil, true)
	if err != nil {
		return err
	}

	data := decodeBodyWithInterface[application.Application](res.Body)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", data.Message)
	}

	runGetApplicationView(&data.Data)

	return nil
}

func runGetApplicationView(app *application.Application) {
	terminal.Header()
	terminal.Gap()

	terminal.ContentKeyValue("ID", app.ID)
	terminal.ContentKeyValue("Name", app.Name)
	terminal.ContentKeyValue("Context", app.Context)
	terminal.ContentKeyValue("Active assignments", fmt.Sprintf("%d", len(app.Assignments)))
	terminal.ContentKeyValue("Created at", app.CreatedAt.Format("2006/01/02"))

	terminal.Gap()
	terminal.Content("Assignments: ", "white")
	terminal.Gap()

	for _, a := range app.Assignments {
		isDone := a.Status == "done"
		statusMarker := helpers.If[string](isDone, "[x]", "[ ]")

		if !isDone {
			terminal.Content(fmt.Sprintf("\t%s %s", statusMarker, a.Title), "danger")
			continue
		}

		terminal.Content(fmt.Sprintf("\t%s %s", statusMarker, a.Title), "success")
	}

	terminal.Gap()
	terminal.Footer()
}
