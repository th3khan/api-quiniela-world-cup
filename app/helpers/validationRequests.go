package helpers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(ctx *fiber.Ctx, target interface{}) error {
	if err := ctx.BodyParser(target); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(target); err != nil {
		return err
	}

	return nil
}

func GetParamFromRequest(ctx *fiber.Ctx, paramName string) string {
	params := ctx.AllParams()
	return params[paramName]
}
