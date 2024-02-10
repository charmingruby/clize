package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/gin-gonic/gin"
)

func NewFetchApplicationsHandler(svc *application.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apps, err := svc.FetchApplication()
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		ctx.JSON(http.StatusOK, apps)
	}
}
