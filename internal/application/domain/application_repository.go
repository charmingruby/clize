package domain

type ApplicationRepository interface {
	Create(application *Application) error
	FindByName(name string) (*Application, error)
	Delete(name string) error
}
