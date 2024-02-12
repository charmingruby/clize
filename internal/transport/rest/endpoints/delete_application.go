package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/gin-gonic/gin"
)

type deleteApplicationResponse struct {
	Message string `json:"message"`
}

func NewDeleteApplicationHandler(svc *application.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		err := svc.DeleteApplication(name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		successMsg := fmt.Sprintf("%s deleted successfully", name)
		res := &deleteApplicationResponse{
			Message: successMsg,
		}

		ctx.JSON(http.StatusOK, res)
	}
}
