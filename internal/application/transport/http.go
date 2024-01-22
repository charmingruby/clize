package transport

import (
	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/charmingruby/clize/internal/application/endpoints"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(r *gin.Engine, svc *domain.Service) *gin.Engine {
	createApplicationHandler := endpoints.NewCreateApplicationHandler(svc.ApplicationService)
	getApplicationHandler := endpoints.NewGetApplicationHandler(svc.ApplicationService)
	deleteApplicationHandler := endpoints.NewDeleteApplicationHandler(svc.ApplicationService)
	fetchApplicationsHandler := endpoints.NewFetchApplicationsHandler(svc.ApplicationService)
	modifyApplicationHandler := endpoints.NewModifyApplicationHandler(svc.ApplicationService)

	r.POST("/applications", createApplicationHandler)
	r.GET("/applications/", fetchApplicationsHandler)
	r.GET("/applications/:name", getApplicationHandler)
	r.PUT("/applications/:name", modifyApplicationHandler)
	r.DELETE("/applications/:name", deleteApplicationHandler)

	return r
}
