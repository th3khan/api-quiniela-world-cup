package auth

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/services/auth"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
)

func Login(ctx *fiber.Ctx) error {
	request := entities.LoginRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(&request)

	if err != nil {
		return err
	}

	err, token, user := auth.Login(request.Email, request.Password)

	if err != nil {
		return err
	}

	var response entities.LoginResponse

	response.User = entities.CreateUserResponse(user)
	response.AccessToken = token

	return ctx.JSON(response)

}
