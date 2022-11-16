package payment

import (
	"image"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func CreatePayment(ctx *fiber.Ctx) error {
	var request entities.PaymentRequest
	err := helpers.ValidateRequest(ctx, &request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var image image.Image
	var imageType string

	image, imageType, err = helpers.ValidateImage(request.Image)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var filenameImage string
	filename := uuid.NewString()
	filenameImage, err = helpers.SaveImageToDisk(image, filename, imageType, models.PATH_FOLDER_PAYMENTS)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "No se pudo guardar la imagen del equipo")
	}

	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	db := database.Connection()
	repo := repositories.NewPaymentRepository(db)

	err, userModel := helpers.GetUserLogged(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	err, payment := repo.CreatePayment(
		date,
		request.Amount,
		filenameImage,
		userModel.ID,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(entities.CreatePaymentResponse(&payment))
}
