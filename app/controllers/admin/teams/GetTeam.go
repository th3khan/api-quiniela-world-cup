package teams

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func GetTeam(ctx *fiber.Ctx) error {
	id, err := validateIdParam(ctx)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID no válido")
	}

	db := database.Connection()
	teamRepository := repositories.NewTeamRepository(db)

	team := teamRepository.GetTeam(uint(id))

	if team.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Equipo no existe")
	}

	return ctx.JSON(entities.CreateTeamResponse(team))
}
