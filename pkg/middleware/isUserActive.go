package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
)

func IsUserActive(c *fiber.Ctx) error {

	err, userModel := helpers.GetUserLogged(c)

	if err != nil {
		return err
	}

	if userModel.Active {
		return c.Next()
	}
	return fiber.NewError(fiber.StatusUnauthorized, "Usuario Inactivo.")
}
