package domain

import "time"

func NewAssignment(title, description, createdBy string) *Assignment {
	a := &Assignment{
		Title:       title,
		Description: description,
		CreatedBy:   createdBy,
		CreateAt:    time.Now(),
		SignedBy:    "",
		SolvedAt:    nil,
	}

	return a
}

type Assignment struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedBy   string     `json:"created_by"`
	SignedBy    string     `json:"signed_by"`
	CreateAt    time.Time  `json:"create_at"`
	SolvedAt    *time.Time `json:"solved_at"`
}

func (a *Assignment) Validate() error {
	return nil
}
