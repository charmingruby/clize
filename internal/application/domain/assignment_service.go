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

func (as *AssignmentService) FetchAssignment() ([]*Assignment, error) {
	assignments, err := as.repo.Fetch()
	if err != nil {
		return nil, err
	}

	return assignments, nil
}

func (as *AssignmentService) FetchAssignmentByApplication(appName string) ([]Assignment, error) {
	assignments, err := as.repo.FetchByApplication(appName)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}
