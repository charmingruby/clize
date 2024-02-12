package rest

import (
	"github.com/charmingruby/clize/internal/domain"
	"github.com/charmingruby/clize/internal/transport/rest/endpoints"
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
	modifyAssignmentHandler := endpoints.NewModifyAssignmentHandler(svc.AssignmentService)
	submitAssignmentHandler := endpoints.NewSubmitAssignmentHandler(svc.AssignmentService)

	signUpHandler := endpoints.NewSignUpHandler(svc.ProfileService)
	signInHandler := endpoints.NewSignInHandler(svc.ProfileService)

	r.POST("/sign-in", signInHandler)
	r.POST("/sign-up", signUpHandler)

	r.POST("/applications", AuthMiddleware(), createApplicationHandler)
	r.GET("/applications/", AuthMiddleware(), fetchApplicationsHandler)
	r.GET("/applications/:name", AuthMiddleware(), getApplicationHandler)
	r.PUT("/applications/:name", AuthMiddleware(), modifyApplicationHandler)
	r.DELETE("/applications/:name", AuthMiddleware(), deleteApplicationHandler)

	r.POST("/applications/:application-name/assignments", AuthMiddleware(), addAssignmentHandler)
	r.PUT("/submit/:application-name/:assignment-id", AuthMiddleware(), submitAssignmentHandler)
	r.PUT("/assignments/:application-name/:assignment-id", AuthMiddleware(), modifyAssignmentHandler)
	r.GET("/applications/assignments/:application-name", AuthMiddleware(), fetchAssignmentsByAppHandler)
	r.GET("/assignments", AuthMiddleware(), fetchAssignmentsHandler)
	r.DELETE("/assignments/:application-name/:assignment-title", AuthMiddleware(), removeAssignmentHandler)

	return r
}
