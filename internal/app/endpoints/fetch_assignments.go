package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/app/domain/assignment"
	"github.com/gin-gonic/gin"
)

func NewFetchAssignmentsHandler(svc *assignment.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		assignments, err := svc.FetchAssignment()
		if err != nil {
			ctx.Status(http.StatusBadRequest)

			return
		}

		ctx.JSON(http.StatusOK, assignments)
	}
}
