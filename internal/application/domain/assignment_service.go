package domain

type AssignmentService struct {
	repo AssignmentRepository
}

func NewAssignmentService(repo AssignmentRepository) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}

func (as *AssignmentService) AddAssignment(applicationName string, assignment *Assignment) error {
	return nil
}
