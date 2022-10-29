package users

import "github.com/gofiber/fiber/v2"

func GetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"users": "lista de usuarios",
	})
}
