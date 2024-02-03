package application

import (
	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/pkg/errors"
)

type ApplicationService struct {
	repo ApplicationRepository
}

func NewApplicationService(repo ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		repo: repo,
	}
}

func (as *ApplicationService) CreateApplication(name, context string) (*Application, error) {
	if _, err := as.repo.FindByName(name); err == nil {
		return nil, &errors.UniqueValueViolationError{
			Field:   "name",
			Entity:  "application",
			Message: errors.NewUniqueValueViolationErrorMessage("application", "name"),
		}
	}

	app, err := NewApplication(name, context)
	if err != nil {
		return nil, err
	}

	if err := as.repo.Create(app); err != nil {
		return nil, err
	}

	return app, nil
}

func (as *ApplicationService) GetApplication(name string) (*Application, error) {
	app, err := as.repo.FindByName(name)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (as *ApplicationService) FetchApplication() ([]*Application, error) {
	app, err := as.repo.Fetch()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (as *ApplicationService) DeleteApplication(name string) error {
	app, err := as.repo.FindByName(name)
	if err != nil {
		return err
	}

	if err := as.repo.Delete(app.Name); err != nil {
		return err
	}

	return nil
}

func (as *ApplicationService) ModifyApplication(identifier, name, context, status string) error {
	app, err := as.repo.FindByName(identifier)
	if err != nil {
		return err
	}

	if app.Name == name {
		as.repo.Delete(app.Name)
	}

	newApp := &Application{
		Name:    helpers.If[string](name == "", app.Name, name),
		Context: helpers.If[string](context == "", app.Context, context),
		Status:  helpers.If[string](status == "", app.Status, status),
	}

	if status != "" {
		if err := newApp.UpdateStatus(status); err != nil {
			return err
		}
	}

	if err := newApp.Validate(); err != nil {
		return err
	}

	if err := as.repo.Create(newApp); err != nil {
		return err
	}

	return nil
}
