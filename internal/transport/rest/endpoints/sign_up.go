package endpoints

import (
	"net/http"

	"github.com/charmingruby/clize/internal/domain/profile"
	"github.com/charmingruby/clize/internal/validation"
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
			errMsg := validation.NewInvalidPayloadErrorMessage(signUpRequiredFields)
			res := WrapResponse[validation.InvalidPayloadError](
				&validation.InvalidPayloadError{
					RequiredFields: signUpRequiredFields,
				},
				http.StatusBadRequest,
				errMsg,
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		if err := svc.Register(
			req.Username,
			req.Email,
			req.Password,
		); err != nil {
			uvv, ok := err.(*validation.UniqueValueViolationError)
			if ok {
				res := WrapResponse[validation.UniqueValueViolationError](
					uvv,
					http.StatusBadRequest,
					uvv.Error(),
				)

				ctx.JSON(http.StatusBadRequest, res)
				return
			}

			res := WrapResponse[error](
				&err,
				http.StatusBadRequest,
				err.Error(),
			)

			ctx.JSON(http.StatusBadRequest, res)
			return
		}

		res := WrapResponse[string](
			nil,
			http.StatusCreated,
			NewCreatedResponse(req.Username),
		)

		ctx.JSON(http.StatusCreated, res)
	}
}
