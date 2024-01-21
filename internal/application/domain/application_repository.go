package domain

type ApplicationRepository interface {
	Create(application *Application) error
	FindByName(name string) (*Application, error)
	FindByKey(key string) (*Application, error)
	Fetch() ([]*Application, error)
	Delete(name string) error
}
