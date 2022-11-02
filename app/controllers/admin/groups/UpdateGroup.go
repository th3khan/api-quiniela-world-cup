package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/app/validations"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func UpdateGroup(ctx *fiber.Ctx) error {
	var err error

	err, id := validations.ValidateIdParam(ctx)
	if err != nil {
		return err
	}

	err, request := validations.ValidateRequest(ctx)
	if err != nil {
		return err
	}

	db := database.Connection()
	repo := repositories.NewGroupRepository(db)

	group := repo.GetGroupById(id)

	if group.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Grupo no existe")
	}

	groupByName := repo.GetGroupByName(request.Name, int(group.ID))
	if groupByName.ID > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Grupo ya se encuentra registrado.")
	}

	err, _ = repo.UpdateGroup(id, request.Name, request.Active)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	group = repo.GetGroupById(id)

	return ctx.JSON(entities.CreateGroupresponse(group))
}
