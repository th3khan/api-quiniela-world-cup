package teams

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func DeleteTeam(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(helpers.GetParamFromRequest(ctx, "id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID no válido")
	}

	db := database.Connection()
	teamRepository := repositories.NewTeamRepository(db)

	err = teamRepository.DeleteTeam(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
