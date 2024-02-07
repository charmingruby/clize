package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/app/domain/application"
	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/fatih/color"
)

func FetchAssignments() error {
	res, err := doRequest(http.MethodGet, "/assignments", nil, true)
	if err != nil {
		cliui.PrintServerError()
		return err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var assignments []application.Assignment
	err = json.Unmarshal(data, &assignments)
	if err != nil {
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

	sort.Slice(assignments, func(i, j int) bool {
		if assignments[i].Status == "done" && assignments[j].Status != "done" {
			return true
		}
		if assignments[i].Status != "done" && assignments[j].Status == "done" {
			return false
		}
		return assignments[i].CreateAt.Before(assignments[j].CreateAt)
	})

	for idx, a := range assignments {
		isAssignmentDone := a.Status == "done"

		status := helpers.If[string](isAssignmentDone, "[x]", "[ ]")

		if isAssignmentDone {
			print(cliui.Padding())
			color.Green("%d. %s %s: %s (%s)", idx+1, status, a.ID, a.Title, a.CreateAt.Format("2006/01/02"))
			amountOfAssignmentsDone++
			continue
		}

		print(cliui.Padding())

		color.Red("%d. %s %s: %s (%s)", idx+1, status, a.ID, a.Title, a.CreateAt.Format("2006/01/02"))

	}
	percentageOfAssignmentsDone := (float64(amountOfAssignmentsDone) / float64(totalAssignments)) * 100

	println(cliui.Gap())
	println(cliui.Content(fmt.Sprintf("%d of %d is done (%.2f%%)", amountOfAssignmentsDone, totalAssignments, percentageOfAssignmentsDone)))
	println(cliui.Footer())
}
