package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type submitAssignmentResponse struct {
	Message string `json:"message"`
}

func NewSubmitAssignmentHandler(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")
		assignmentId := ctx.Param("assignment-id")

		if err := svc.SubmitAssignment(appName, assignmentId); err != nil {
			rnf, ok := err.(*errors.ResourceNotFoundError)
			if ok {
				ctx.JSON(http.StatusNotFound, rnf)
				return
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		res := &submitAssignmentResponse{
			Message: fmt.Sprintf("%s: %s submitted successfully", appName, assignmentId),
		}

		ctx.JSON(http.StatusOK, res)
	}
}
