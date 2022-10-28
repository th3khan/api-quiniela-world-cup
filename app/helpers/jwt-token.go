package helpers

import (
	"github.com/dgrijalva/jwt-go"
)

var secretKey = "1714FA0476C8CA25C53FFD74FBBFD8DE3BC7ED7FDBA79B47FE05A597A000381F"

func GenerateToken(claims jwt.Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return webtoken, nil
}
