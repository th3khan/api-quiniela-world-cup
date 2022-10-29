package users

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities/admin"
)

func CreateUser(ctx *fiber.Ctx) error {
	var request admin.CreateUserRequest

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(&request)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"usuario": "crear",
	})
}
