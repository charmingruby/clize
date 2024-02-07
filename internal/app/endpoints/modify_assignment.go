package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type modifyAssignmentRequest struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

var modifyAssignmentFieldsOptions = []string{
	"name", "context", "status", "solved at",
}

func NewModifyAssignmentHandler(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("assignment-id")

		var req modifyAssignmentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			err = &errors.InvalidPayloadError{
				RequiredFields: modifyAppFieldsOptions,
				Message:        errors.NewInvalidPayloadErrorMessage(modifyAssignmentFieldsOptions),
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if req.Context == "" && req.Name == "" {
			err := &errors.NotNullableBodyError{
				Message: errors.NewNotNullableErrorMessage(modifyAppFieldsOptions),
				Fields:  modifyAssignmentFieldsOptions,
			}
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := svc.UpdateAssignment(id, req.Name, req.Context); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
