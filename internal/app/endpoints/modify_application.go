package endpoints

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type modifyApplicationRequest struct {
	Name    string `json:"name"`
	Context string `json:"context"`
	Status  string `json:"status"`
}

type modifyApplicationResponse struct {
	Message string `json:"message"`
}

var modifyAppFieldsOptions = []string{
	"name", "context", "status",
}

func NewModifyApplicationHandler(svc *application.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		var req modifyApplicationRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if req.Name == "" && req.Context == "" && req.Status == "" {
			err := errors.NotNullableBodyError{
				Message: errors.NewNotNullableErrorMessage(modifyAppFieldsOptions),
				Fields:  modifyAppFieldsOptions,
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := svc.ModifyApplication(name, req.Name, req.Context, req.Status); err != nil {
			errType := reflect.TypeOf(err)

			if errType.Name() == "ResourceNotFoundError" {
				err := errors.ResourceNotFoundError{
					Entity:  "application",
					Message: errors.NewResourceNotFoundErrorMessage("application"),
				}

				ctx.JSON(http.StatusNotFound, err)
				return
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		successMsg := fmt.Sprintf("%s modified successfully", name)
		res := &modifyApplicationResponse{
			Message: successMsg,
		}

		ctx.JSON(http.StatusOK, res)
	}
}
