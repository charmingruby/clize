package rest

import (
	"encoding/gob"

	"github.com/charmingruby/clize/internal/auth"
	"github.com/charmingruby/clize/internal/transport/rest/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New(auth *auth.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/login", handlers.LoginHandler(auth))
	router.GET("/callback", handlers.CallbackHandler(auth))
	router.GET("/user", handlers.UserProfileHandler)

	return router
}
