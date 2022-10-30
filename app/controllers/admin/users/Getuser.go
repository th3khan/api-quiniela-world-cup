package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/services/admin"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
)

func GetUser(ctx *fiber.Ctx) error {
	params := ctx.AllParams()

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Id no valido")
	}

	user := admin.GetUserById(id)

	if user.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Usuario no existe")
	}

	return ctx.JSON(entities.CreateUserResponse(&user))
}
