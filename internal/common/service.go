package common

import (
	"github.com/charmingruby/clize/internal/common/http/transport/rest"
	"github.com/gin-gonic/gin"
)

func NewHTTPService(r *gin.Engine) (*gin.Engine, error) {
	r = rest.NewHTTPHandler(r)

	return r, nil
}
