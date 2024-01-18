package domain

type ApplicationService struct {
	repo ApplicationRepository
}

func NewApplicationService(repo ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		repo: repo,
	}
}

func (as *ApplicationService) CreateApplication(name, context string) (*Application, error) {
	app, err := NewApplication(name, context)
	if err != nil {
		return nil, err
	}

	if err := as.repo.Create(app); err != nil {
		return nil, err
	}

	return app, nil
}
