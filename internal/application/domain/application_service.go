package domain

import "github.com/charmingruby/clize/pkg/errors"

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
