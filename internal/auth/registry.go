package auth

import (
	"context"

	"github.com/charmingruby/clize/internal/auth/domain"
	"github.com/charmingruby/clize/internal/auth/repository/redis"
	"github.com/charmingruby/clize/internal/auth/transport/rest"
	"github.com/gin-gonic/gin"
	rdb "github.com/go-redis/redis/v8"
)

func NewService(rc *rdb.Client) (*domain.Service, error) {
	ctx := context.Background()

	// Instantiate repos
	profileRepo := redis.NewRedisProfileRepository(ctx, rc)

	// Instantiate services
	profileSvc := domain.NewProfileService(profileRepo)

	// Centralize services
	svc := &domain.Service{
		ProfileService: profileSvc,
	}

	return svc, nil
}

func NewHTTPService(r *gin.Engine, svc *domain.Service) (*gin.Engine, error) {
	r = rest.NewHTTPHandler(r, svc)

	return r, nil
}
