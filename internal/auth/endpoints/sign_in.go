package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/auth/domain"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/charmingruby/clize/pkg/token"
	"github.com/gin-gonic/gin"
)

type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var signInRequiredFields = []string{"username", "password"}

func NewSignInHandler(svc *domain.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req signInRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			err = &errors.InvalidPayloadError{
				Message:        errors.NewInvalidPayloadErrorMessage(signInRequiredFields),
				RequiredFields: signInRequiredFields,
			}
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := svc.ProfileService.Login(req.Username, req.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		token, err := token.NewJwtService().GenerateToken(req.Username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
