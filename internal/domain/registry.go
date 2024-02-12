package domain

import (
	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/domain/profile"
)

type Service struct {
	ApplicationService *application.ApplicationService
	AssignmentService  *application.AssignmentService
	ProfileService     *profile.ProfileService
}
