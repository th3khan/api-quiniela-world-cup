package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
)

func IsSuperAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	roleUser := claims["roleId"].(float64)

	fmt.Println()

	if int(roleUser) == models.ROLE_SUPER_ADMIN {
		return c.Next()
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Usuario sin privilegios.")
}
