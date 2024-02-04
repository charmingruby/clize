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

	r.POST("/applications", common.AuthMiddleware(), createApplicationHandler)
	r.GET("/applications/", common.AuthMiddleware(), fetchApplicationsHandler)
	r.GET("/applications/:name", common.AuthMiddleware(), getApplicationHandler)
	r.PUT("/applications/:name", common.AuthMiddleware(), modifyApplicationHandler)
	r.DELETE("/applications/:name", common.AuthMiddleware(), deleteApplicationHandler)

	r.POST("/applications/:application-name/assignments", common.AuthMiddleware(), addAssignmentHandler)
	r.GET("/applications/assignments/:application-name", common.AuthMiddleware(), fetchAssignmentsByAppHandler)
	r.GET("/assignments", common.AuthMiddleware(), fetchAssignmentsHandler)
	r.DELETE("/assignments/:application-name/:assignment-title", common.AuthMiddleware(), removeAssignmentHandler)

	return r
}
