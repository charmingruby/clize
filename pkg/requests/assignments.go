package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/app/domain/application"
	cliui "github.com/charmingruby/clize/pkg/cli_ui"
)

func FetchAssignments() error {
	res, err := doRequest(http.MethodGet, "/assignments", nil, true)
	if err != nil {
		println("erro na request")
		return err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println("erro no body")

		return err
	}

	var assignments []application.Assignment
	err = json.Unmarshal(data, &assignments)
	if err != nil {
		println("erro no unmarshal")

		return err
	}

	tableView(assignments)

	return nil
}

func tableView(assignments []application.Assignment) {
	var amountOfAssignmentsDone int
	totalAssignments := len(assignments)

	println(cliui.Header())
	println(cliui.Title("Total Assignments"))
	println(cliui.Gap())

	for idx, a := range assignments {
		isAssignmentDone := a.Status == "done"

		if isAssignmentDone {
			amountOfAssignmentsDone++
		}

		status := helpers.If[string](isAssignmentDone, "[x]", "[ ]")

		println(cliui.Content(fmt.Sprintf("%d. %s %s: %s (%s)", idx+1, status, a.ID, a.Title, a.CreateAt.Format("2006/01/02"))))
	}
	percentageOfAssignmentsDone := (amountOfAssignmentsDone / totalAssignments) * 100

	println(cliui.Gap())
	println(cliui.Content(fmt.Sprintf("%d of %d is done (%d%%)", amountOfAssignmentsDone, totalAssignments, percentageOfAssignmentsDone)))
	println(cliui.Footer())
}
