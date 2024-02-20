package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type removeAssignmentResponse struct {
	Message string `json:"message"`
}

func NewRemoveAssignmentHandler(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")
		assignmentName := ctx.Param("assignment-title")

		if err := svc.RemoveAssignmentFromApplication(
			appName,
			assignmentName,
		); err != nil {

			rnf, ok := err.(*errors.ResourceNotFoundError)
			if ok {
				ctx.JSON(http.StatusNotFound, rnf)
				return
			}

			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		successMsg := fmt.Sprintf("%s deleted successfully from %s", assignmentName, appName)
		res := &removeAssignmentResponse{
			Message: successMsg,
		}
		ctx.JSON(http.StatusOK, res)
	}
}
