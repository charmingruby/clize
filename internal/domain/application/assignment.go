package application

import (
	"strings"
	"time"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/internal/validation"
	"github.com/charmingruby/clize/pkg/uuid"
)

func NewAssignment(title, description string) (*Assignment, error) {
	sts := Status()

	formattedTitle := formatTitle(title)

	a := &Assignment{
		ID:          uuid.GenerateUUID(),
		Title:       formattedTitle,
		Description: description,
		Status:      sts["awaiting"],
		CreatedAt:   time.Now(),
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
	CreatedAt   time.Time  `json:"create_at"`
	SolvedAt    *time.Time `json:"solved_at,omitempty"`
}

func (a *Assignment) Validate() error {
	if a.Title == "" {
		return &validation.RequiredFieldError{
			FieldName: "title",
		}
	}

	if len(a.Title) < 4 {
		return &validation.FieldLengthError{
			IsMinimumError: true,
			Quantity:       4,
			FieldName:      "title",
		}
	}

	if len(a.Title) > 20 {
		return &validation.FieldLengthError{
			IsMinimumError: false,
			Quantity:       20,
			FieldName:      "title",
		}
	}

	if a.Description == "" {
		return &validation.RequiredFieldError{
			FieldName: "description",
		}
	}
	if len(a.Description) < 8 {

		return &validation.FieldLengthError{
			IsMinimumError: true,
			Quantity:       8,
			FieldName:      "description",
		}
	}

	if len(a.Description) > 100 {
		return &validation.FieldLengthError{
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

func (a *Assignment) Submit() {
	sts := Status()
	now := time.Now()

	a.Status = sts["done"]
	a.SolvedAt = &now
}

func formatTitle(title string) string {
	return strings.ReplaceAll(title, " ", "_")
}
