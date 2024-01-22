package domain

type AssignmentService struct {
	repo AssignmentRepository
}

func NewAssignmentService(repo AssignmentRepository) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}
