package application

import (
	"time"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/charmingruby/clize/pkg/uuid"
)

func NewApplication(name, context string) (*Application, error) {
	sts := Status()

	a := &Application{
		ID:          uuid.GenerateUUID(),
		Name:        name,
		Status:      sts["awaiting"],
		Context:     context,
		Assignments: []Assignment{},
		CreatedAt:   time.Now(),
	}

	err := a.Validate()
	if err != nil {
		return nil, err
	}

	return a, nil
}

type Application struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Context     string       `json:"context"`
	Status      string       `json:"status"`
	Assignments []Assignment `json:"assignments"`
	CreatedAt   time.Time    `json:"created_at"`
}

func (a *Application) UpdateStatus(status string) error {
	a.Status = status

	if err := ValidateStatus(a.Status); err != nil {
		return err
	}

	return nil
}

func (a *Application) Validate() error {
	if a.Name == "" {
		return &errors.RequiredFieldError{FieldName: "name"}
	}

	if len(a.Name) < 4 {
		return &errors.FieldLengthError{
			IsMinimumError: true,
			Quantity:       4,
			FieldName:      "name",
		}
	}

	if len(a.Name) > 24 {
		return &errors.FieldLengthError{
			IsMinimumError: false,
			Quantity:       24,
			FieldName:      "name",
		}
	}

	if a.Context == "" {
		return &errors.RequiredFieldError{FieldName: "context"}
	}

	if len(a.Context) < 6 {
		return &errors.FieldLengthError{
			IsMinimumError: true,
			Quantity:       6,
			FieldName:      "context",
		}
	}

	if len(a.Context) > 40 {
		return &errors.FieldLengthError{
			IsMinimumError: false,
			Quantity:       40,
			FieldName:      "context",
		}
	}

	if err := ValidateStatus(a.Status); err != nil {
		return err
	}

	return nil
}

func (a *Application) Modify(name, context string) error {
	a.Name = helpers.If[string](name == "", a.Name, name)
	a.Context = helpers.If[string](context == "", a.Context, context)
	return a.Validate()
}

func (a *Application) ProgressReview() {
	sts := Status()

	if len(a.Assignments) == 0 {
		a.Status = sts["done"]
		return
	}

	a.Status = sts["awaiting"]
}

func (a *Application) SetAssignments(assignments []Assignment) {
	a.Assignments = assignments
}
