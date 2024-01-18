package domain

type ApplicationRepository interface {
	Create(application *Application) error
	FindByIdentifier(identifier string) *Application
	Delete(identifier string) error
}
