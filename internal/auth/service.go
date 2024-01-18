package auth

import (
	"github.com/charmingruby/clize/internal/auth/domain"
	"github.com/charmingruby/clize/internal/auth/transport"
	"github.com/gin-gonic/gin"
)

func NewHTTPService(r *gin.Engine, a *domain.Authenticator) (*gin.Engine, error) {

	r = transport.NewHTTPHandler(r, a)

	return r, nil
}
