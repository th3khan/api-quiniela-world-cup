package teams

import (
	"image"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func UpdateTeam(ctx *fiber.Ctx) error {
	id, err := validateIdParam(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID no válido.")
	}

	request, err := validateRequest(ctx)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	db := database.Connection()
	teamRepository := repositories.NewTeamRepository(db)

	team := teamRepository.GetTeam(uint(id))
	if team.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Equipo no existe")
	}

	var image image.Image
	var imageType string
	var filenameImage string

	if team.Logo != request.Logo {
		image, imageType, err = helpers.ValidateImage(request.Logo)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		filename := uuid.NewString()
		filenameImage, err = helpers.SaveImageToDisk(image, filename, imageType, models.PATH_FOLDER_LOGO_TEAMS)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "No se pudo guardar la imagen del equipo")
		}
	} else {
		filenameImage = request.Logo
	}

	err, team = teamRepository.UpdateTeam(
		uint(id),
		request.Name,
		request.Active,
		filenameImage,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(entities.CreateTeamResponse(team))
}
