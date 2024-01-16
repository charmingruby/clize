package apps

type AppService struct {
	repo AppRepository
}

func NewAppService(repo AppRepository) *AppService {
	return &AppService{
		repo: repo,
	}
}

func (as *AppService) CreateApp(app *App) error {

	as.repo.Create(app)

	return nil
}
