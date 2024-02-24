package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type modifyAssignmentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var modifyAssignmentFieldsOptions = []string{
	"title", "description",
}

func NewModifyAssignmentHandler(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		applicationName := ctx.Param("application-name")
		assignmentTitle := ctx.Param("assignment-title")

		var req modifyAssignmentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			err = &errors.InvalidPayloadError{
				RequiredFields: modifyAppFieldsOptions,
				Message:        errors.NewInvalidPayloadErrorMessage(modifyAssignmentFieldsOptions),
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if req.Title == "" && req.Description == "" {
			err := &errors.NotNullableBodyError{
				Message: errors.NewNotNullableErrorMessage(modifyAssignmentFieldsOptions),
				Fields:  modifyAssignmentFieldsOptions,
			}
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := svc.UpdateAssignment(assignmentTitle, applicationName, req.Title, req.Description); err != nil {
			rnf, ok := err.(*errors.ResourceNotFoundError)
			if ok {
				ctx.JSON(http.StatusNotFound, rnf)
				return
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		successMsg := fmt.Sprintf("%s modified successfully", assignmentTitle)
		res := &modifyApplicationResponse{
			Message: successMsg,
		}
		ctx.JSON(http.StatusOK, res)
	}
}
