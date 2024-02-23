package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/clize/internal/domain/profile"
	"github.com/charmingruby/clize/pkg/errors"
	"github.com/gin-gonic/gin"
)

type signUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var signUpRequiredFields = []string{"username", "email", "password"}

func NewSignUpHandler(svc *profile.ProfileService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req signUpRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			err = &errors.InvalidPayloadError{
				Message:        errors.NewInvalidPayloadErrorMessage(signUpRequiredFields),
				RequiredFields: signUpRequiredFields,
			}
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if err := svc.Register(
			req.Username,
			req.Email,
			req.Password,
		); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			NewCreatedResponse(fmt.Sprintf("%s's %s", req.Username, "profile")),
		)
	}
}