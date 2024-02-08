package application

import (
	"strings"
	"time"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/charmingruby/clize/pkg/uuid"
)

func NewAssignment(title, description, createdBy string) (*Assignment, error) {
	sts := Status()

	formattedTitle := formatTitle(title)

	a := &Assignment{
		ID:          uuid.GenerateUUID(),
		Title:       formattedTitle,
		Description: description,
		CreatedBy:   createdBy,
		Status:      sts["awaiting"],
		CreateAt:    time.Now(),
		SignedBy:    "",
		SolvedAt:    nil,
	}

	if err := a.Validate(); err != nil {
		return nil, err
	}

	return a, nil
}

type Assignment struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedBy   string     `json:"created_by"`
	SignedBy    string     `json:"signed_by,omitempty"`
	CreateAt    time.Time  `json:"create_at"`
	SolvedAt    *time.Time `json:"solved_at,omitempty"`
}

func (a *Assignment) Validate() error {
	if a.Title == "" {
		return &errors.RequiredFieldError{
			FieldName: "title",
		}
	}

	if len(a.Title) < 4 {
		return &errors.FieldLengthError{
			IsMinimumError: true,
			Quantity:       4,
			FieldName:      "title",
		}
	}

	if len(a.Title) > 20 {
		return &errors.FieldLengthError{
			IsMinimumError: false,
			Quantity:       20,
			FieldName:      "title",
		}
	}

	if a.Description == "" {
		return &errors.RequiredFieldError{
			FieldName: "description",
		}
	}
	if len(a.Description) < 8 {

		return &errors.FieldLengthError{
			IsMinimumError: true,
			Quantity:       8,
			FieldName:      "description",
		}
	}

	if len(a.Description) > 100 {
		return &errors.FieldLengthError{
			IsMinimumError: false,
			Quantity:       100,
			FieldName:      "description",
		}
	}

	if err := ValidateStatus(a.Status); err != nil {
		return err
	}

	return nil
}

func (a *Assignment) Modify(title, description string) error {
	a.Title = helpers.If[string](title == "", a.Title, title)
	a.Description = helpers.If[string](description == "", a.Description, description)
	return a.Validate()
}

func formatTitle(title string) string {
	return strings.ReplaceAll(title, " ", "_")
}
