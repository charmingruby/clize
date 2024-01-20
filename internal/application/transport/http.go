package transport

import (
	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/charmingruby/clize/internal/application/endpoints"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(r *gin.Engine, svc *domain.ApplicationService) *gin.Engine {
	createApplicationHandler := endpoints.NewCreateApplicationHandler(svc)
	getApplicationHandler := endpoints.NewGetApplicationHandler(svc)
	deleteApplicationHandler := endpoints.NewDeleteApplicationHandler(svc)

	r.POST("/application", createApplicationHandler)
	r.GET("/application/:name", getApplicationHandler)
	r.DELETE("/application/:name", deleteApplicationHandler)

	return r
}
