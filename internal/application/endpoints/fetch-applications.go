package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/application/domain"
	"github.com/gin-gonic/gin"
)

type fetchApplicationResponse struct {
	Applications []*domain.Application `json:"applications"`
}

func NewFetchApplicationsHandler(svc *domain.ApplicationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apps, err := svc.FetchApplication()
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}

		res := &fetchApplicationResponse{
			Applications: apps,
		}

		ctx.JSON(http.StatusOK, res)
	}
}
