package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
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
			errMsg := validation.NewInvalidPayloadErrorMessage(modifyAppFieldsOptions)
			res := WrapResponse[validation.InvalidPayloadError](
				&validation.InvalidPayloadError{
					RequiredFields: modifyAppFieldsOptions,
				},
				http.StatusBadRequest,
				errMsg,
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		if req.Title == "" && req.Description == "" {
			errMsg := validation.NewNotNullableErrorMessage(modifyAssignmentFieldsOptions)
			res := WrapResponse[validation.NotNullableBodyError](
				&validation.NotNullableBodyError{
					Fields: modifyAssignmentFieldsOptions,
				},
				http.StatusBadRequest,
				errMsg,
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		if err := svc.UpdateAssignment(assignmentTitle, applicationName, req.Title, req.Description); err != nil {
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

			res := WrapResponse[error](
				&err,
				http.StatusNotFound,
				err.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := WrapResponse[string](
			nil,
			http.StatusOK,
			NewModifiedResponse(assignmentTitle),
		)

		ctx.JSON(http.StatusOK, res)
	}
}
