package helpers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/th3khan/api-quiniela-world-cup/config"
)

func GenerateToken(claims jwt.Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(config.JWTSecret))

	if err != nil {
		return "", err
	}

	return webtoken, nil
}
