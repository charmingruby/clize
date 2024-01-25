package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(r *gin.Engine) *gin.Engine {
	r.GET("health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"is-healthy": true,
		})
	})

	return r
}
