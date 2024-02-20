package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

func NewFetchAssignmentsByApplication(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")

		assignments, err := svc.FetchAssignmentByApplication(appName)
		if err != nil {
			rnf, ok := err.(*errors.ResourceNotFoundError)
			if ok {
				ctx.JSON(http.StatusNotFound, rnf)
				return
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, assignments)
	}
}
