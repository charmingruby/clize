package internal

import (
	"context"

	"github.com/charmingruby/clize/internal/database/redis"
	"github.com/charmingruby/clize/internal/domain"
	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/domain/profile"
	"github.com/charmingruby/clize/internal/transport/rest"
	"github.com/gin-gonic/gin"
	rdb "github.com/go-redis/redis/v8"
)

func NewService(rc *rdb.Client) (*domain.Service, error) {
	ctx := context.Background()

	// Instantiate repos
	applicationRepo := redis.NewRedisApplicationRepository(rc, ctx)
	profileRepo := redis.NewRedisProfileRepository(ctx, rc)

	// Instatiate services
	applicationService := application.NewApplicationService(applicationRepo)
	assignmentService := application.NewAssignmentService(applicationRepo)
	profileService := profile.NewProfileService(profileRepo)

	// Centralize services
	svc := &domain.Service{
		ApplicationService: applicationService,
		AssignmentService:  assignmentService,
		ProfileService:     profileService,
	}

	return svc, nil
}

func NewHTTPService(r *gin.Engine, svc *domain.Service) (*gin.Engine, error) {
	r = rest.NewHTTPHandler(r, svc)
	return r, nil
}
