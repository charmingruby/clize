package assignment

import (
	"time"

	"github.com/charmingruby/clize/internal/app/domain/common"
	"github.com/charmingruby/clize/pkg/errors"
)

func NewAssignment(title, description string, createdBy int) (*Assignment, error) {
	sts := common.Status()

	a := &Assignment{
		Title:       title,
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
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedBy   int        `json:"created_by"`
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

	if err := common.ValidateStatus(a.Status); err != nil {
		return err
	}

	return nil
}

func (a *Assignment) Sign(githubUsername string) {
	now := time.Now()
	nowPointer := &now

	a.SignedBy = githubUsername
	a.SolvedAt = nowPointer
}
