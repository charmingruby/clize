package endpoints

import (
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
		assignmentId := ctx.Param("assignment-id")
		applicationName := ctx.Param("application-name")

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

		if err := svc.UpdateAssignment(assignmentId, applicationName, req.Title, req.Description); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
