package groups

import (
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/app/validations"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func CreateGroup(ctx *fiber.Ctx) error {
	err, request := validations.ValidateRequest(ctx)
	if err != nil {
		return err
	}

	db := database.Connection()
	repo := repositories.NewGroupRepository(db)

	groupByName := repo.GetGroupByName(request.Name, 0)
	if groupByName.ID > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Grupo ya se encuentra registrado.")
	}

	err, newGroup := repo.CreateGroup(request.Name, request.Active)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(entities.CreateGroupresponse(newGroup))
}
