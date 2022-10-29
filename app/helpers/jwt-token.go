package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/services/admin"
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

func GetUserLogged(c *fiber.Ctx) (error, models.User) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	email := claims["email"].(string)

	err, userModel := admin.GetUserByEmail(email)

	if err != nil {
		return err, models.User{}
	}

	return nil, userModel
}
