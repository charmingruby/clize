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
	fetchApplicationsHandler := endpoints.NewFetchApplicationsHandler(svc)

	r.POST("/applications", createApplicationHandler)
	r.GET("/applications/", fetchApplicationsHandler)
	r.GET("/applications/:name", getApplicationHandler)
	r.DELETE("/applications/:name", deleteApplicationHandler)

	return r
}
