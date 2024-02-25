package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
	"github.com/gin-gonic/gin"
)

func NewFetchAssignmentsByApplication(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")

		assignments, err := svc.FetchAssignmentByApplication(appName)
		if err != nil {
			rnf, ok := err.(*validation.ResourceNotFoundError)
			if ok {
				res := WrapResponse[validation.ResourceNotFoundError](
					rnf,
					http.StatusNotFound,
					rnf.Error(),
				)

				ctx.JSON(http.StatusNotFound, res)
				return
			}

			res := WrapResponse[string](
				nil,
				http.StatusBadRequest,
				err.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		ctx.JSON(http.StatusOK, assignments)
	}
}
