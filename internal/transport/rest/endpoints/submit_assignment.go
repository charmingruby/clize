package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"

	"github.com/gin-gonic/gin"
)

func NewSubmitAssignmentHandler(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")
		assignmentTitle := ctx.Param("assignment-title")

		if err := svc.SubmitAssignment(appName, assignmentTitle); err != nil {
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

			res := WrapResponse[error](
				&err,
				http.StatusBadRequest,
				err.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		msg := fmt.Sprintf("\"%s\": \"%s\" submitted successfully", appName, assignmentTitle)

		res := WrapResponse[string](
			nil,
			http.StatusOK,
			msg,
		)

		ctx.JSON(http.StatusOK, res)
	}
}
