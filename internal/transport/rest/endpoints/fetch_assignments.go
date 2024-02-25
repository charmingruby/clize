package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/gin-gonic/gin"
)

func NewFetchAssignmentsHandler(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		assignments, err := svc.FetchAssignment()
		if err != nil {
			res := WrapResponse[error](
				&err,
				http.StatusBadRequest,
				err.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := WrapResponse[[]application.Assignment](
			&assignments,
			http.StatusOK,
			NewFetchedResponse("assignment", len(assignments)),
		)

		ctx.JSON(http.StatusOK, res)
	}
}
