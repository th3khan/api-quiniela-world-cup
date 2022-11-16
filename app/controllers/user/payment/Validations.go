package payment

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
)

func validateRequest(ctx *fiber.Ctx) (entities.PaymentRequest, error) {
	var request entities.PaymentRequest
	if err := ctx.BodyParser(&request); err != nil {
		return request, err
	}

	validate := validator.New()
	if err := validate.Struct(&request); err != nil {
		return request, err
	}

	return request, nil
}
