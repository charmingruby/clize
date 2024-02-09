package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/app/domain/application"
	cliui "github.com/charmingruby/clize/pkg/cli_ui"
	"github.com/fatih/color"
)

type fetchAssignmentsOutput struct {
	Assignments []application.Assignment `json:"assignments"`
}

func FetchAssignments() error {
	res, err := doRequest(http.MethodGet, "/assignments", nil, true)
	if err != nil {
		cliui.PrintServerError()
		return err
	}

	op, err := decodeFetchAssignments(res.Body)
	if err != nil {
		return err
	}

	runFetchAssignmentsView(op.Assignments)

	return nil
}

func decodeFetchAssignments(body io.ReadCloser) (*fetchAssignmentsOutput, error) {
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var assignments []application.Assignment
	err = json.Unmarshal(data, &assignments)
	if err != nil {
		return nil, err
	}

	return &fetchAssignmentsOutput{
		Assignments: assignments,
	}, nil
}

func runFetchAssignmentsView(assignments []application.Assignment) {
	var amountOfAssignmentsDone int
	totalAssignments := len(assignments)

	cliui.Header()
	cliui.Gap()
	cliui.Title("Total Assignments")
	cliui.Gap()

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
			cliui.Padding()
			color.Green("%d. %s %s: %s (%s)", idx+1, status, a.ID, a.Title, a.CreateAt.Format("2006/01/02"))

			amountOfAssignmentsDone++
			continue
		}

		cliui.Padding()
		color.Red("%d. %s %s: %s (%s)", idx+1, status, a.ID, a.Title, a.CreateAt.Format("2006/01/02"))

	}
	percentageOfAssignmentsDone := (float64(amountOfAssignmentsDone) / float64(totalAssignments)) * 100

	cliui.Gap()
	cliui.Content(fmt.Sprintf("%d of %d is done (%.2f%%)", amountOfAssignmentsDone, totalAssignments, percentageOfAssignmentsDone))
	cliui.Gap()
	cliui.Footer()
}
