package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/app/validations"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func GetGroup(ctx *fiber.Ctx) error {
	err, id := validations.ValidateIdParam(ctx)
	if err != nil {
		return err
	}

	db := database.Connection()
	repo := repositories.NewGroupRepository(db)

	group := repo.GetGroupById(id)

	if group.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Grupo no existe")
	}

	return ctx.JSON(entities.CreateGroupresponse(group))
}
