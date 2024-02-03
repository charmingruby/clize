package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/app/domain/assignment"
	"github.com/gin-gonic/gin"
)

func NewFetchAssignmentsByApplication(svc *assignment.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")

		assgnmnts, err := svc.FetchAssignmentByApplication(appName)
		if err != nil {
			ctx.Status(http.StatusNotFound)
		}

		ctx.JSON(http.StatusOK, assgnmnts)
	}
}
