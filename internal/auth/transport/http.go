package transport

import (
	"encoding/gob"

	"github.com/charmingruby/clize/internal/auth/domain"
	"github.com/charmingruby/clize/internal/auth/endpoints"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewHTTPHandler(r *gin.Engine, a *domain.Authenticator) *gin.Engine {

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("auth-session", store))

	r.GET("/login", endpoints.NewLoginHandler(a))
	r.GET("/callback", endpoints.NewCallbackHandler(a))
	r.GET("/user", endpoints.NewUserProfileHandler)

	return r
}
