package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHealthCheckHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	}
}
