package common

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPService(r *gin.Engine) (*gin.Engine, error) {
	r = NewHTTPHandler(r)

	return r, nil
}
