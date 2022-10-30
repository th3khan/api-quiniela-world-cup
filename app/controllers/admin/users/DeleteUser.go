package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/services/admin"
)

func DeleteUser(ctx *fiber.Ctx) error {
	params := ctx.AllParams()

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Id no valido")
	}

	user := admin.GetUserById(id)

	if user.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Usuario no existe")
	}

	if err = admin.DeleteUserById(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "No se pudo eliminar el usuario")
	}

	return ctx.Status(fiber.StatusNoContent).SendString("")
}
