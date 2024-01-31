package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type createApplicationRequest struct {
	Name    string `json:"name"`
	Context string `json:"context"`
}

var createApplicationRequiredFields = []string{"name", "context"}

type createApplicationResponse struct {
	Message string `json:"message"`
}

func NewCreateApplicationHandler(svc *domain.ApplicationService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var req createApplicationRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			err = &errors.InvalidPayloadError{
				Message:        errors.NewInvalidPayloadErrorMessage(createApplicationRequiredFields),
				RequiredFields: createApplicationRequiredFields,
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		app, err := svc.CreateApplication(req.Name, req.Context)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		successMsg := fmt.Sprintf("%s created successfully", app.Name)
		res := &createApplicationResponse{
			Message: successMsg,
		}

		ctx.JSON(http.StatusCreated, res)
	}
}
