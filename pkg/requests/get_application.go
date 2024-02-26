package requests

import (
	"fmt"
	"net/http"

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
