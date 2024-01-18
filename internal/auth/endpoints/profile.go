package endpoints

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func NewUserProfileHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	ctx.JSON(200, profile)
}
