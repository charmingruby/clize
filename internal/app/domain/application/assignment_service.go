package application

import (
	"github.com/charmingruby/clize/pkg/errors"
)

type AssignmentService struct {
	repo ApplicationRepository
}

func NewAssignmentService(repo ApplicationRepository) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}

func (as *AssignmentService) AddAssignment(applicationName, title, description, createdBy string) error {
	newAssignment, err := NewAssignment(
		title, description, createdBy,
	)
	if err != nil {
		return err
	}

	app, err := as.repo.FindByName(applicationName)
	if err != nil {
		return err
	}

	app.Assignments = append(app.Assignments, *newAssignment)

	return as.repo.Create(app)
}

func (as *AssignmentService) FetchAssignment() ([]Assignment, error) {
	apps, err := as.repo.Fetch()
	if err != nil {
		return nil, err
	}

	var assignments []Assignment
	for _, app := range apps {
		assignments = append(assignments, app.Assignments...)
	}

	return assignments, nil
}

func (as *AssignmentService) FetchAssignmentByApplication(appName string) ([]Assignment, error) {
	app, err := as.repo.FindByName(appName)
	if err != nil {
		return nil, err
	}

	return app.Assignments, nil
}

func (as *AssignmentService) RemoveAssignmentFromApplication(appName, assignmentName string) error {
	app, err := as.repo.FindByName(appName)
	if err != nil {
		return err
	}

	newAssignments := []Assignment{}
	var assignmentFound bool = false
	for _, assignment := range app.Assignments {
		if assignment.Title == assignmentName {
			assignmentFound = true
			continue
		}

		newAssignments = append(newAssignments, assignment)
	}

	if !assignmentFound {
		return &errors.ResourceNotFoundError{
			Entity:  "assignment",
			Message: errors.NewResourceNotFoundErrorMessage("assignment"),
		}
	}

	app.SetAssignments(newAssignments)

	err = as.repo.Create(app)
	if err != nil {
		return err
	}

	return nil
}

func (as *AssignmentService) UpdateAssignment(id, applicationName, title, description string) error {
	app, err := as.repo.FindByName(applicationName)
	if err != nil {
		return &errors.ResourceNotFoundError{
			Entity:  "application",
			Message: errors.NewResourceNotFoundErrorMessage("application"),
		}
	}

	var assignmentToModify *Assignment
	var assignments []Assignment
	for _, a := range app.Assignments {
		if a.ID == id {
			assignmentToModify = &a
			continue
		}

		assignments = append(assignments, a)
	}

	if assignmentToModify == nil {
		return &errors.ResourceNotFoundError{
			Entity:  "assignment",
			Message: errors.NewResourceNotFoundErrorMessage("assignment"),
		}
	}

	if err := assignmentToModify.Modify(title, description); err != nil {
		return err
	}

	assignments = append(assignments, *assignmentToModify)
	app.Assignments = assignments

	as.repo.Create(app)

	return nil
}
