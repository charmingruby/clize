package rest

import (
	"github.com/charmingruby/clize/internal/auth/domain"
	"github.com/charmingruby/clize/internal/auth/endpoints"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(r *gin.Engine, svc *domain.Service) *gin.Engine {
	signUpHandler := endpoints.NewSignUpHandler(svc)
	signInHandler := endpoints.NewSignInHandler(svc)

	r.POST("/sign-in", signInHandler)
	r.POST("/sign-up", signUpHandler)

	return r
}
