package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/gin-gonic/gin"
)

type removeAssignmentResponse struct {
	Message string `json:"message"`
}

func NewRemoveAssignmentHandler(svc *domain.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appName := ctx.Param("application-name")
		assignmentName := ctx.Param("assignment_name")

		if err := svc.RemoveAssignmentFromApplication(
			appName,
			assignmentName,
		); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		successMsg := fmt.Sprintf("%s deleted successfully from %s", assignmentName, appName)
		res := &removeAssignmentResponse{
			Message: successMsg,
		}
		ctx.JSON(http.StatusOK, res)
	}
}
