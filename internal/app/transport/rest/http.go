package rest

import (
	"github.com/charmingruby/clize/internal/app/domain"
	"github.com/charmingruby/clize/internal/app/endpoints"
	"github.com/charmingruby/clize/internal/common"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(r *gin.Engine, svc *domain.Service) *gin.Engine {
	createApplicationHandler := endpoints.NewCreateApplicationHandler(svc.ApplicationService)
	getApplicationHandler := endpoints.NewGetApplicationHandler(svc.ApplicationService)
	deleteApplicationHandler := endpoints.NewDeleteApplicationHandler(svc.ApplicationService)
	fetchApplicationsHandler := endpoints.NewFetchApplicationsHandler(svc.ApplicationService)
	modifyApplicationHandler := endpoints.NewModifyApplicationHandler(svc.ApplicationService)

	addAssignmentHandler := endpoints.NewAddAssignmentHandler(svc.AssignmentService)
	fetchAssignmentsHandler := endpoints.NewFetchAssignmentsHandler(svc.AssignmentService)
	fetchAssignmentsByAppHandler := endpoints.NewFetchAssignmentsByApplication(svc.AssignmentService)
	removeAssignmentHandler := endpoints.NewRemoveAssignmentHandler(svc.AssignmentService)

	private := r.Group("admin", common.AuthMiddleware())
	{
		private.POST("/applications", createApplicationHandler)
		private.GET("/applications/", fetchApplicationsHandler)
		private.GET("/applications/:name", getApplicationHandler)
		private.PUT("/applications/:name", modifyApplicationHandler)
		private.DELETE("/applications/:name", deleteApplicationHandler)
		private.POST("/applications/:application-name/assignments", addAssignmentHandler)
		private.GET("/applications/assignments/:application-name", fetchAssignmentsByAppHandler)
		private.GET("/assignments", fetchAssignmentsHandler)
		private.DELETE("/assignments/:application-name/:assignment-title", removeAssignmentHandler)
	}

	return r
}
