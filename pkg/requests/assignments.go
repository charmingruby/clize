package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/app/domain/application"
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

	println("~# C l i z e")
	println("~  ( Total Assignments )")
	println("~")

	for _, a := range assignments {
		isAssignmentDone := a.Status == "done"

		if isAssignmentDone {
			amountOfAssignmentsDone++
		}

		status := helpers.If[string](isAssignmentDone, "[x]", "[ ]")
		fmt.Printf("~  %s %s: %s (%s) \n", status, a.ID, a.Title, a.CreateAt.Format("2006/01/02"))
	}
	percentageOfAssignmentsDone := (amountOfAssignmentsDone / totalAssignments) * 100

	println("~")
	fmt.Printf("~  %d of %d is done (%d%%)\n", amountOfAssignmentsDone, totalAssignments, percentageOfAssignmentsDone)
	println("~#")
}
