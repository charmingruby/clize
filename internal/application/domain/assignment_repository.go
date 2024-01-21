package domain

type AssignmentRepository interface {
	Create(assignment *Assignment) error
	FindByApplicationName(applicationName string) (*Assignment, error)
	Sign(assignment *Assignment) error
	Delete(assignment *Assignment) error
	Modify(assignment *Assignment) error
}
