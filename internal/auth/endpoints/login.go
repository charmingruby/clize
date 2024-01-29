package endpoints

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/charmingruby/clize/internal/auth/domain"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginResponse struct {
	RedirectUrl string `json:"redirect_url"`
}

func NewLoginHandler(auth *domain.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the state inside the session.
		session := sessions.Default(ctx)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		res := loginResponse{
			RedirectUrl: auth.AuthCodeURL(state),
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
