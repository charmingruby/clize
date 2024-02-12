package rest

import (
	"github.com/charmingruby/clize/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const bearerPattern = "Bearer "

		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(401)
		}

		t := header[len(bearerPattern):]

		if isTokenValid := token.NewJwtService().ValidateToken(t); !isTokenValid {
			ctx.AbortWithStatus(401)
		}
	}
}
