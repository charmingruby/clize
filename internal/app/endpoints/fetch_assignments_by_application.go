package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/gin-gonic/gin"
)

func NewFetchAssignmentsByApplication(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")

		assignments, err := svc.FetchAssignmentByApplication(appName)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		ctx.JSON(http.StatusOK, assignments)
	}
}
