package application

import "github.com/charmingruby/clize/internal/validation"

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
		return nil, &validation.UniqueValueViolationError{
			Field:   "name",
			Entity:  "application",
			Message: validation.NewUniqueValueViolationErrorMessage("application", "name"),
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

func (as *ApplicationService) ModifyApplication(identifier, name, context string) error {
	app, err := as.repo.FindByName(identifier)
	if err != nil {
		return err
	}

	if app.Name == name {
		as.repo.Delete(app.Name)
	}

	if err := app.Modify(name, context); err != nil {
		return err
	}

	if err := as.repo.Delete(identifier); err != nil {
		return err
	}

	as.repo.Create(app)

	return nil
}
