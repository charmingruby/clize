package rest

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(r *gin.Engine) *gin.Engine {
	r.POST("/sign-in")
	r.POST("/sign-up")

	return r
}
