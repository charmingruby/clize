package domain

type AssignmentService struct {
	repo AssignmentRepository
}

func NewAssignmentService(repo AssignmentRepository) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}

func (as *AssignmentService) AddAssignment(applicationName, title, description string, createdBy int) error {
	newAssignment, err := NewAssignment(
		title, description, createdBy,
	)
	if err != nil {
		return err
	}

	if err := as.repo.CreateAndAddToApplication(applicationName, newAssignment); err != nil {
		return err
	}

	return nil
}
