package requests

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/terminal"
)

func FetchAssignments() error {
	res, err := doRequest(http.MethodGet, "/assignments", nil, true)
	if err != nil {
		terminal.PrintServerError()
		return err
	}

	data := decodeBodyWithInterface[[]application.Assignment](res.Body)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", data.Message)
	}

	runFetchAssignmentsView(data.Data)

	return nil
}

func runFetchAssignmentsView(assignments []application.Assignment) {
	var amountOfAssignmentsDone int
	totalAssignments := len(assignments)

	terminal.Header()
	terminal.Gap()

	if totalAssignments == 0 {
		terminal.Padding()
		//terminal.boldGreen.Printf("No assignments.\n")

		terminal.Gap()
		terminal.Footer()
		return
	}

	terminal.Title("Total Assignments")
	terminal.Gap()

	sortedAssignments := sortAssignments(assignments)

	for idx, a := range sortedAssignments {
		isAssignmentDone := a.Status == "done"

		status := helpers.If[string](isAssignmentDone, "[x]", "[ ]")

		if isAssignmentDone {
			terminal.Content(
				fmt.Sprintf("%d. %s %s (%s)", idx+1, status, a.Title, a.CreatedAt.Format("2006/01/02")),
				"lsuccess",
			)
			terminal.Content(
				fmt.Sprintf("\t -> %s", a.Description),
				"",
			)

			amountOfAssignmentsDone++
			terminal.Gap()
			continue
		}

		terminal.Content(
			fmt.Sprintf("%d. %s %s (%s)", idx+1, status, a.Title, a.CreatedAt.Format("2006/01/02")),
			"ldanger",
		)
		terminal.Content(
			fmt.Sprintf("\t -> %s", a.Description),
			"",
		)
		terminal.Gap()
	}

	percentageOfAssignmentsDone := (float64(amountOfAssignmentsDone) / float64(totalAssignments)) * 100

	terminal.Content(fmt.Sprintf("%d of %d is done (%.2f%%)", amountOfAssignmentsDone, totalAssignments, percentageOfAssignmentsDone), "")
	terminal.Gap()
	terminal.Footer()
}

func sortAssignments(assignments []application.Assignment) []application.Assignment {
	sort.Slice(assignments, func(i, j int) bool {
		if assignments[i].Status == "done" && assignments[j].Status != "done" {
			return true
		}
		if assignments[i].Status != "done" && assignments[j].Status == "done" {
			return false
		}
		return assignments[i].CreatedAt.Before(assignments[j].CreatedAt)
	})

	return assignments
}
