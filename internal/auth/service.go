package auth

import (
	"github.com/charmingruby/clize/internal/auth/domain"
	"github.com/charmingruby/clize/internal/auth/transport/rest"
	"github.com/gin-gonic/gin"
)

func NewHTTPService(r *gin.Engine, a *domain.Authenticator) (*gin.Engine, error) {

	r = rest.NewHTTPHandler(r, a)

	return r, nil
}
