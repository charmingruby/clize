package domain

import (
	"time"

	"github.com/charmingruby/clize/pkg/errors"
)

func NewAssignment(title, description, createdBy string) *Assignment {
	sts := Status()

	a := &Assignment{
		Title:       title,
		Description: description,
		CreatedBy:   createdBy,
		Status:      sts[awaiting],
		CreateAt:    time.Now(),
		SignedBy:    "",
		SolvedAt:    nil,
	}

	return a
}

type Assignment struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedBy   string     `json:"created_by"`
	SignedBy    string     `json:"signed_by"`
	CreateAt    time.Time  `json:"create_at"`
	SolvedAt    *time.Time `json:"solved_at"`
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

	if len(a.Title) > 32 {
		return &errors.FieldLengthError{
			IsMinimumError: false,
			Quantity:       32,
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

	if len(a.Description) > 48 {
		return &errors.FieldLengthError{
			IsMinimumError: false,
			Quantity:       48,
			FieldName:      "description",
		}
	}

	if err := ValidateStatus(a.Status); err != nil {
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
