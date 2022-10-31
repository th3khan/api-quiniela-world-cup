package groups

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func UpdateGroup(ctx *fiber.Ctx) error {
	var request entities.GroupRequest
	var err error
	params := ctx.AllParams()

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Id no valido")
	}

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(&request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	db := database.Connection()
	repo := repositories.NewGroupRepository(db)

	group := repo.GetGroupById(id)

	if group.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Grupo no existe")
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
