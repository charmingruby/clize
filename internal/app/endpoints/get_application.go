package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/gin-gonic/gin"
)

type getApplicationResponse struct {
	Application *application.Application `json:"application"`
}

func NewGetApplicationHandler(svc *application.ApplicationService) gin.HandlerFunc {
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
