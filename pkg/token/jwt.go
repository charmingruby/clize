package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretKey: "secret",
		issuer:    "clize",
	}
}

type JwtClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (js *jwtService) GenerateToken(username string) (string, error) {
	expirationTime := time.Duration(time.Minute * 60 * 24)

	claim := &JwtClaim{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(expirationTime).Unix(),
			Issuer:    js.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(js.secretKey))
	if err != nil {
		return "", nil
	}

	return t, nil
}

func (js *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(js.secretKey), nil
	})

	return err == nil
}

func RetriveSubValues(token string) {}
