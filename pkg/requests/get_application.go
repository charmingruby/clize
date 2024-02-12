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

type getApplicationInput struct {
	Name string `json:"name"`
}

type getApplicationOutput struct {
	Application application.Application `json:"application"`
}

func GetApplication(name string) error {
	inputs := &getApplicationInput{
		Name: name,
	}

	res, err := doRequest(http.MethodGet, fmt.Sprintf("/applications/%s", inputs.Name), nil, true)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusNotFound {
		terminal.PrintNotFoundResponse(name)
		return err
	}

	op, err := decodeGetApplicationBody(res.Body)
	if err != nil {
		return err
	}

	runGetApplicationView(&op.Application)

	return nil
}

func decodeGetApplicationBody(body io.ReadCloser) (*getApplicationOutput, error) {
	defer body.Close()
	result, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var parsedResponse application.Application
	err = json.Unmarshal(result, &parsedResponse)
	if err != nil {
		return nil, err
	}

	return &getApplicationOutput{
		Application: parsedResponse,
	}, nil
}

func runGetApplicationView(app *application.Application) {
	terminal.Header()
	terminal.Gap()
	terminal.Title(app.Name)
	terminal.Gap()

	terminal.Content(fmt.Sprintf("ID: %s", app.ID))
	terminal.Content(fmt.Sprintf("Context: %s", app.Context))
	terminal.Content(fmt.Sprintf("Status: %s", app.Status))
	terminal.Content(fmt.Sprintf("CreatedAt: %s", app.CreatedAt.Format("2006/01/02")))

	terminal.Content("Assignments: ")
	terminal.Content("[")

	for _, a := range app.Assignments {
		terminal.Padding()
		terminal.Tab(1)
		fmt.Println("{")

		terminal.Padding()
		terminal.Tab(2)
		fmt.Printf("ID: %s\n", a.ID)

		terminal.Padding()
		terminal.Tab(2)
		fmt.Printf("Title: %s\n", a.Title)

		terminal.Padding()
		terminal.Tab(2)
		fmt.Printf("Description: %s\n", a.Description)

		terminal.Padding()
		terminal.Tab(2)
		fmt.Printf("Status: %s\n", a.Status)

		terminal.Padding()
		terminal.Tab(2)
		fmt.Printf("CreatedAt: %s\n", a.CreatedAt.Format("2006/01/02"))

		terminal.Padding()
		terminal.Tab(1)
		fmt.Println("},")
	}

	terminal.Content("]")

	terminal.Gap()
	terminal.Footer()
}
