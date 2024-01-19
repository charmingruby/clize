package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/gin-gonic/gin"
)

type getApplicationResponse struct {
	Application *domain.Application `json:"application"`
}

func NewGetApplicationHandler(svc *domain.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		app, err := svc.GetApplication(name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		res := &getApplicationResponse{
			Application: app,
		}

		ctx.JSON(http.StatusOK, res)
	}
}
