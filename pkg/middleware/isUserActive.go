package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/th3khan/api-quiniela-world-cup/app/services/admin"
)

func IsUserActive(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	email := claims["email"].(string)

	err, userModel := admin.GetUserByEmail(email)

	if err != nil {
		return err
	}

	if userModel.Active {
		return c.Next()
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Usuario Inactivo.")
}
