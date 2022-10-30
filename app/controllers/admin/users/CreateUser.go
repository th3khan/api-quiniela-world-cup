package users

import (
	"image"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities/admin"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func CreateUser(ctx *fiber.Ctx) error {
	var request admin.CreateUserRequest

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(&request)

	if err != nil {
		return err
	}

	var image image.Image
	var imageType string
	imageIsDefinedAndValid := false

	if len(request.Image) > 0 {
		image, imageType, err = helpers.ValidateImage(request.Image)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		imageIsDefinedAndValid = true
	}

	db := database.Connection()
	userRepository := repositories.NewUserRespository(db)
	roleRepository := repositories.NewRoleRepository(db)

	err, _ = roleRepository.GetRoleById(request.RoleId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Rol no existe")
	}

	userByEmail := userRepository.GetUserByEmail(request.Email)

	if userByEmail.ID > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Usuario ya existe.")
	}

	var filenameImage string
	if imageIsDefinedAndValid {
		filename := uuid.NewString()
		filenameImage, err = helpers.SaveImageToDisk(image, filename, imageType, models.PATH_PROFILE_IMAGES)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "No se pudo guardar la imagen del usuario")
		}

	}

	err, newUser := userRepository.CreateUser(
		request.Name,
		request.Email,
		request.RoleId,
		request.Password,
		request.Active,
		filenameImage,
		request.EmailVerified,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "No se pudo crear el usuario")
	}

	return ctx.Status(fiber.StatusCreated).JSON(entities.CreateUserResponse(&newUser))
}
