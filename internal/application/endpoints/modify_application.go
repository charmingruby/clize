package endpoints

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/charmingruby/clize/internal/application/domain"
	cErrors "github.com/charmingruby/clize/pkg/errors"
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

var fieldsOptions = []string{
	"name", "context", "status",
}

func NewModifyApplicationHandler(svc *domain.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		var req modifyApplicationRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if req.Name == "" && req.Context == "" && req.Status == "" {

			err := cErrors.NotNullableBodyError{
				Message: cErrors.NewNotNullableErrorMessage(fieldsOptions),
				Fields:  fieldsOptions,
			}

			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := svc.ModifyApplication(name, req.Name, req.Context, req.Status); err != nil {
			errType := reflect.TypeOf(err)

			if errType.Name() == "ResourceNotFoundError" {
				err := cErrors.ResourceNotFoundError{
					Entity:  "application",
					Message: cErrors.NewResourceNotFoundErrorMessage("application"),
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
