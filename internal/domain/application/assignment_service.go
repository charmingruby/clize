package application

import "github.com/charmingruby/clize/internal/validation"

type AssignmentService struct {
	repo ApplicationRepository
}

func NewAssignmentService(repo ApplicationRepository) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}

func (as *AssignmentService) AddAssignment(applicationName, title, description string) error {
	newAssignment, err := NewAssignment(
		title, description,
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
		return &validation.ResourceNotFoundError{
			Entity: "assignment",
		}
	}

	app.SetAssignments(newAssignments)

	err = as.repo.Create(app)
	if err != nil {
		return err
	}

	return nil
}

func (as *AssignmentService) UpdateAssignment(assignmentTitle, applicationName, title, description string) error {
	app, err := as.repo.FindByName(applicationName)
	if err != nil {
		return &validation.ResourceNotFoundError{
			Entity: "application",
		}
	}

	var assignmentToModify *Assignment
	var assignments []Assignment
	for _, a := range app.Assignments {
		if a.Title == assignmentTitle {
			assignmentToModify = &a
			continue
		}

		assignments = append(assignments, a)
	}

	if assignmentToModify == nil {
		return &validation.ResourceNotFoundError{
			Entity: "assignment",
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

func (as *AssignmentService) SubmitAssignment(applicationName, assignmentTitle string) error {
	app, err := as.repo.FindByName(applicationName)
	if err != nil {
		return &validation.ResourceNotFoundError{
			Entity: "application",
		}
	}

	var assignmentToSubmit *Assignment
	var assignments []Assignment
	for _, a := range app.Assignments {
		if a.Title == assignmentTitle {
			assignmentToSubmit = &a
			continue
		}

		assignments = append(assignments, a)
	}

	if assignmentToSubmit == nil {
		return &validation.ResourceNotFoundError{
			Entity: "assignment",
		}
	}

	assignmentToSubmit.Submit()

	assignments = append(assignments, *assignmentToSubmit)
	app.Assignments = assignments

	as.repo.Create(app)

	return nil
}
