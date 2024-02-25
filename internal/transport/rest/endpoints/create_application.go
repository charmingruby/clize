package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
	"github.com/gin-gonic/gin"
)

type createApplicationRequest struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

var createApplicationRequiredFields = []string{"name", "context"}

func NewCreateApplicationHandler(svc *application.ApplicationService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req createApplicationRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errMsg := validation.NewInvalidPayloadErrorMessage(createApplicationRequiredFields)
			res := WrapResponse[validation.InvalidPayloadError](
				&validation.InvalidPayloadError{
					RequiredFields: createApplicationRequiredFields,
				},
				http.StatusBadRequest,
				errMsg,
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		_, err := svc.CreateApplication(req.Name, req.Context)
		if err != nil {
			res := WrapResponse[error](
				&err,
				http.StatusBadRequest,
				err.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := WrapResponse[string](
			nil,
			http.StatusCreated,
			NewCreatedResponse(req.Name),
		)

		ctx.JSON(http.StatusCreated, res)
	}
}
