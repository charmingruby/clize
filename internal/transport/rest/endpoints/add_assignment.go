package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type addAssignmentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var addAssignmentRequiredFields = []string{"title", "description"}

func NewAddAssignmentHandler(svc *application.AssignmentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		applicationName := ctx.Param("application-name")

		var req addAssignmentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			err = &errors.InvalidPayloadError{
				Message:        errors.NewInvalidPayloadErrorMessage(addAssignmentRequiredFields),
				RequiredFields: addAssignmentRequiredFields,
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		// TODO: handle with session value
		createdBy := "static profile"

		if err := svc.AddAssignment(applicationName, req.Title, req.Description, createdBy); err != nil {
			rnf, ok := err.(*errors.ResourceNotFoundError)
			if ok {
				ctx.JSON(http.StatusNotFound, rnf)
				return
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusCreated, NewCreatedResponse(req.Title))
	}
}