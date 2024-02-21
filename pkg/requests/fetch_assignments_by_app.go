package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/terminal"
	"github.com/fatih/color"
)

type fetchAssignmentsByApplicationOutput struct {
	Assignments []application.Assignment `json:"assignments"`
}

func FetchAssignmentsByApplication(appName string) error {
	url := fmt.Sprintf("/applications/assignments/%s", appName)

	res, err := doRequest(http.MethodGet, url, nil, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	statusCode := res.StatusCode
	if statusCode != http.StatusOK {
		if statusCode == http.StatusNotFound {
			errRes := decodeNotFoundError(res.Body)
			terminal.PrintErrorResponse(errRes.Message)
			return err
		}
	}

	op, err := decodeFetchAssignmentsByAppBody(res.Body)
	if err != nil {
		return err
	}

	runFetchAssignmentsByAppView(appName, op.Assignments)

	return nil
}

func runFetchAssignmentsByAppView(appName string, assignments []application.Assignment) {
	terminal.ClearTerminal()

	var amountOfAssignmentsDone int
	totalAssignments := len(assignments)

	terminal.Header()
	terminal.Gap()

	if totalAssignments == 0 {
		terminal.Padding()
		terminal.BoldGreen.Printf("No assignments.\n")

		terminal.Gap()
		terminal.Footer()
		return
	}

	terminal.Title(fmt.Sprintf("%s's Assignments", appName))
	terminal.Gap()

	sort.Slice(assignments, func(i, j int) bool {
		if assignments[i].Status == "done" && assignments[j].Status != "done" {
			return true
		}
		if assignments[i].Status != "done" && assignments[j].Status == "done" {
			return false
		}
		return assignments[i].CreatedAt.Before(assignments[j].CreatedAt)
	})

	for idx, a := range assignments {
		isAssignmentDone := a.Status == "done"

		status := helpers.If[string](isAssignmentDone, "[x]", "[ ]")

		if isAssignmentDone {
			terminal.Padding()
			color.Green("%d. %s %s: %s (%s)", idx+1, status, a.ID, a.Title, a.CreatedAt.Format("2006/01/02"))

			amountOfAssignmentsDone++
			continue
		}

		terminal.Padding()
		color.Red("%d. %s %s: %s (%s)", idx+1, status, a.ID, a.Title, a.CreatedAt.Format("2006/01/02"))

	}
	percentageOfAssignmentsDone := (float64(amountOfAssignmentsDone) / float64(totalAssignments)) * 100

	terminal.Gap()
	terminal.Content(fmt.Sprintf("%d of %d is done (%.2f%%)", amountOfAssignmentsDone, totalAssignments, percentageOfAssignmentsDone))
	terminal.Gap()
	terminal.Footer()
}

func decodeFetchAssignmentsByAppBody(body io.ReadCloser) (*fetchAssignmentsByApplicationOutput, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var assignments []application.Assignment
	err = json.Unmarshal(data, &assignments)
	if err != nil {
		return nil, err
	}

	return &fetchAssignmentsByApplicationOutput{
		Assignments: assignments,
	}, nil
}
