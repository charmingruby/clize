package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
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
			errMsg := validation.NewInvalidPayloadErrorMessage(addAssignmentRequiredFields)
			res := WrapResponse[validation.InvalidPayloadError](
				&validation.InvalidPayloadError{
					RequiredFields: addAssignmentRequiredFields,
				},
				http.StatusBadRequest,
				errMsg,
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		if err := svc.AddAssignment(applicationName, req.Title, req.Description); err != nil {
			rnf, ok := err.(*validation.ResourceNotFoundError)
			if ok {
				res := WrapResponse[validation.ResourceNotFoundError](
					rnf,
					http.StatusNotFound,
					rnf.Error(),
				)

				ctx.JSON(http.StatusNotFound, res)
				return
			}

			res := WrapResponse[string](
				nil,
				http.StatusBadRequest,
				rnf.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := WrapResponse[string](
			nil,
			http.StatusCreated,
			NewAddItemResponse(req.Title, applicationName),
		)

		ctx.JSON(http.StatusCreated, res)
	}
}
