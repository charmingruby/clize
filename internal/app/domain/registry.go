package domain

import (
	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/charmingruby/clize/internal/app/domain/assignment"
)

type Service struct {
	ApplicationService *application.ApplicationService
	AssignmentService  *assignment.AssignmentService
}
