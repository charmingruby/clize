package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
	"github.com/gin-gonic/gin"
)

type modifyApplicationRequest struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

var modifyAppFieldsOptions = []string{
	"name", "context",
}

func NewModifyApplicationHandler(svc *application.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		var req modifyApplicationRequest
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

		if req.Name == "" && req.Context == "" {
			errMsg := validation.NewNotNullableErrorMessage(modifyAppFieldsOptions)
			res := WrapResponse[validation.NotNullableBodyError](
				&validation.NotNullableBodyError{
					Fields: modifyAppFieldsOptions,
				},
				http.StatusBadRequest,
				errMsg,
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		if err := svc.ModifyApplication(name, req.Name, req.Context); err != nil {
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
			NewModifiedResponse(name),
		)

		ctx.JSON(http.StatusOK, res)
	}
}
