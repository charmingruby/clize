package application

import (
	"context"

	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/charmingruby/clize/internal/application/repository/redis"
	"github.com/charmingruby/clize/internal/application/transport"
	"github.com/gin-gonic/gin"
	rdb "github.com/go-redis/redis/v8"
)

func NewService(rc *rdb.Client) (*domain.Service, error) {
	ctx := context.Background()

	// Instantiate repos
	applicationRepo := redis.NewRedisApplicationRepository(rc, ctx)
	assignmentRepo := redis.NewRedisAssignmentRepository(rc, ctx)

	// Instatiate services
	applicationService := domain.NewApplicationService(applicationRepo)
	assignmentService := domain.NewAssignmentService(assignmentRepo)

	// Centralize services
	svc := &domain.Service{
		ApplicationService: applicationService,
		AssignmentService:  assignmentService,
	}

	return svc, nil
}

func NewHTTPService(r *gin.Engine, svc *domain.Service) (*gin.Engine, error) {
	r = transport.NewHTTPHandler(r, svc)
	return r, nil
}
