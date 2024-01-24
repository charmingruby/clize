package domain

type AssignmentRepository interface {
	CreateAndAddToApplication(applicationName string, assignment *Assignment) error
	Fetch() ([]*Assignment, error)
	FetchByApplication(appName string) ([]Assignment, error)
	FindByApplicationName(applicationName string) (*Assignment, error)
	FindByTitle(title string) (*Assignment, error)
	Modify(assignment *Assignment) error
	Sign(assignment *Assignment) error
	Delete(assignment *Assignment) error
}
