package application

import (
	"github.com/charmingruby/clize/internal/application/database/redis"
	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/charmingruby/clize/internal/application/transport"
	"github.com/gin-gonic/gin"
	rdb "github.com/go-redis/redis/v8"
)

func NewService(rc *rdb.Client) (*domain.ApplicationService, error) {
	applicationRepo := redis.NewRedisApplicationRepository(rc)
	service := domain.NewApplicationService(applicationRepo)
	return service, nil
}

func NewHTTPService(r *gin.Engine, svc *domain.ApplicationService) (*gin.Engine, error) {
	r = transport.NewHTTPHandler(r, svc)
	return r, nil
}
