package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type addAssignmentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var addAssignmentRequiredFields = []string{"title", "description"}

func NewAddAssignmentHandler(svc *domain.Service) gin.HandlerFunc {
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
		createdById := 1

		// check if profile exists

		newAssignment, err := domain.NewAssignment(
			req.Title, req.Description, createdById,
		)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := svc.AssignmentService.AddAssignment(applicationName, newAssignment); err != nil {
			//errType := reflect.TypeOf(err)

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
