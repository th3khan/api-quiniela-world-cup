package teams

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
)

func validateRequest(ctx *fiber.Ctx) (entities.TeamBase, error) {
	var request entities.TeamBase
	if err := ctx.BodyParser(&request); err != nil {
		return request, err
	}

	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		return request, err
	}

	return request, nil
}

func validateIdParam(ctx *fiber.Ctx) (int, error) {
	params := ctx.AllParams()
	id, err := strconv.Atoi(params["id"])
	return id, err
}
