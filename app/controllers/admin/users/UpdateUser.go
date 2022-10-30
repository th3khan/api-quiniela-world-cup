package users

import (
	"image"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/app/services/admin"
	"github.com/th3khan/api-quiniela-world-cup/pkg/entities"
	adminrequest "github.com/th3khan/api-quiniela-world-cup/pkg/entities/admin"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func UpdateUser(ctx *fiber.Ctx) error {
	params := ctx.AllParams()

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Id no valido")
	}

	var request adminrequest.UserUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	validate := validator.New()
	err = validate.Struct(&request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
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

	user := admin.GetUserById(id)

	if user.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Usuario no existe")
	}

	passwordChanged := false
	var passwordHashed string

	if len(request.Password) > 0 {
		passwordChanged = true
		passwordHashed, err = helpers.HashingPassword(request.Password)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	db := database.Connection()
	userRepository := repositories.NewUserRespository(db)
	roleRepository := repositories.NewRoleRepository(db)

	err, _ = roleRepository.GetRoleById(request.RoleId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Rol no existe")
	}

	var filenameImage string
	var imageChanged bool
	if imageIsDefinedAndValid {
		filename := uuid.NewString()
		filenameImage, err = helpers.SaveImageToDisk(image, filename, imageType, models.PATH_PROFILE_IMAGES)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "No se pudo guardar la imagen del usuario")
		}
		imageChanged = true
	}

	var setEmailVerifiedNow bool
	if request.EmailVerified != user.EmailVerified {
		if request.EmailVerified {
			setEmailVerifiedNow = true
		} else {
			setEmailVerifiedNow = false
		}
	}

	if err = userRepository.UpdateUser(
		id,
		request.Name,
		request.Email,
		request.RoleId,
		passwordHashed,
		request.Active,
		filenameImage,
		request.EmailVerified,
		passwordChanged,
		imageChanged,
		setEmailVerifiedNow,
	); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error al intentar actualizar usuario")
	}

	user = userRepository.GetUserById(id)

	return ctx.JSON(entities.CreateUserResponse(&user))
}
