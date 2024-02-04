package domain

import (
	"github.com/charmingruby/clize/internal/app/domain/application"
)

type Service struct {
	ApplicationService *application.ApplicationService
	AssignmentService  *application.AssignmentService
}
