package apps

type AppRepository interface {
	Create(app *App) error
}
