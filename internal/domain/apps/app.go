package apps

import (
	"github.com/charmingruby/clize/internal/domain"
)

type App struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

func NewApp(name, context string) (*App, error) {
	a := &App{
		Name:    name,
		Context: context,
	}

	err := a.Validate()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Validate() error {
	if a.Name == "" {
		return &domain.RequiredFieldError{FieldName: "name"}
	}

	if len(a.Name) < 4 {
		return &domain.FieldLengthError{
			IsMinimumError: true,
			Quantity:       4,
			FieldName:      "name",
		}
	}

	if len(a.Name) > 24 {
		return &domain.FieldLengthError{
			IsMinimumError: false,
			Quantity:       24,
			FieldName:      "name",
		}
	}

	if a.Context == "" {
		return &domain.RequiredFieldError{FieldName: "context"}
	}

	if len(a.Context) < 6 {
		return &domain.FieldLengthError{
			IsMinimumError: true,
			Quantity:       6,
			FieldName:      "context",
		}
	}

	if len(a.Context) > 24 {
		return &domain.FieldLengthError{
			IsMinimumError: false,
			Quantity:       24,
			FieldName:      "context",
		}
	}

	return nil
}
