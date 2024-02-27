package requests

import (
	"fmt"
	"net/http"
	"sort"

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
	terminal.Title("Applications")
	terminal.Gap()

	sortedAppsByAssignmentsAmount := sortAppsByAmountOfAssignments(apps)

	for idx, a := range sortedAppsByAssignmentsAmount {
		var amountDone int = 0

		for _, t := range sortedAppsByAssignmentsAmount {
			if t.Status == "Done" {
				amountDone++
			}
		}

		isCompleted := amountDone == len(a.Assignments)

		if !isCompleted {
			terminal.Content(
				fmt.Sprintf("%d. %d/%d - %s: %s  (%s)", idx+1, amountDone, len(a.Assignments), a.Name, makeExcerpt(a.Context), a.CreatedAt.Format("2006/01/02")),
				"ldanger",
			)
		} else {
			terminal.Content(
				fmt.Sprintf("%d. %d/%d - %s: %s  (%s)", idx+1, amountDone, len(a.Assignments), a.Name, makeExcerpt(a.Context), a.CreatedAt.Format("2006/01/02")),
				"lsuccess",
			)
		}
	}

	terminal.Gap()

	terminal.Content(
		fmt.Sprintf("%d applications found", len(apps)),
		"",
	)

	terminal.Gap()

	terminal.Footer()
}

func sortAppsByAmountOfAssignments(apps []application.Application) []application.Application {
	sort.Slice(apps, func(i, j int) bool {
		if len(apps[i].Assignments) >= len(apps[j].Assignments) {
			return true
		}

		return apps[i].CreatedAt.Before(apps[j].CreatedAt)
	})

	return apps
}
