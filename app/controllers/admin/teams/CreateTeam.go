package teams

import (
	"image"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
)

func CreateTeam(ctx *fiber.Ctx) error {
	request, err := validateRequest(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var image image.Image
	var imageType string

	image, imageType, err = helpers.ValidateImage(request.Logo)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var filenameImage string
	filename := uuid.NewString()
	filenameImage, err = helpers.SaveImageToDisk(image, filename, imageType, models.PATH_FOLDER_LOGO_TEAMS)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "No se pudo guardar la imagen del equipo")
	}

	teamRepository := repositories.NewTeamRepository()

	err, team := teamRepository.CreateTeam(
		request.Name,
		request.Active,
		filenameImage,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "No se pudo guardar la imagen del equipo")
	}

	return ctx.Status(fiber.StatusCreated).JSON(entities.CreateTeamResponse(team))
}
