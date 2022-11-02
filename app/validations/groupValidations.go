package validations

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
)

func ValidateRequest(ctx *fiber.Ctx) (error, entities.GroupRequest) {
	var request entities.GroupRequest
	var err error

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error()), request
	}

	validate := validator.New()
	err = validate.Struct(&request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error()), request
	}

	return nil, request
}

func ValidateIdParam(ctx *fiber.Ctx) (error, int) {
	params := ctx.AllParams()
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Id no valido"), id
	}

	return nil, id
}
