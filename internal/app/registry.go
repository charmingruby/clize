package app

import (
	"context"

	"github.com/charmingruby/clize/internal/app/database/redis"
	"github.com/charmingruby/clize/internal/app/domain"
	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/charmingruby/clize/internal/app/domain/assignment"
	"github.com/charmingruby/clize/internal/app/transport/rest"
	"github.com/gin-gonic/gin"
	rdb "github.com/go-redis/redis/v8"
)

func NewService(rc *rdb.Client) (*domain.Service, error) {
	ctx := context.Background()

	// Instantiate repos
	applicationRepo := redis.NewRedisApplicationRepository(rc, ctx)
	assignmentRepo := redis.NewRedisAssignmentRepository(rc, ctx)

	// Instatiate services
	applicationService := application.NewApplicationService(applicationRepo)
	assignmentService := assignment.NewAssignmentService(assignmentRepo)

	// Centralize services
	svc := &domain.Service{
		ApplicationService: applicationService,
		AssignmentService:  assignmentService,
	}

	return svc, nil
}

func NewHTTPService(r *gin.Engine, svc *domain.Service) (*gin.Engine, error) {
	r = rest.NewHTTPHandler(r, svc)
	return r, nil
}
