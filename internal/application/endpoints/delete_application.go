package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/gin-gonic/gin"
)

type deleteApplicationResponse struct {
	Message string `json:"message"`
}

func NewDeleteApplicationHandler(svc *domain.ApplicationService) gin.HandlerFunc {
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
