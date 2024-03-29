package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/profile"
	"github.com/charmingruby/clize/internal/validation"

	"github.com/charmingruby/clize/pkg/token"
	"github.com/gin-gonic/gin"
)

type signInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var signInRequiredFields = []string{"username", "password"}

func NewSignInHandler(svc *profile.ProfileService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req signInRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			errMsg := validation.NewInvalidPayloadErrorMessage(signInRequiredFields)
			res := WrapResponse[validation.InvalidPayloadError](
				&validation.InvalidPayloadError{
					RequiredFields: signInRequiredFields,
				},
				http.StatusBadRequest,
				errMsg,
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		if err := svc.Login(req.Username, req.Password); err != nil {
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
