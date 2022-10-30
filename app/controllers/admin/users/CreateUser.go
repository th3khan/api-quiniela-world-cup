package users

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
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

	err, newUser := userRepository.CreateUser(
		request.Name,
		request.Email,
		request.RoleId,
		request.Password,
		request.Active,
	)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "No se pudo crear el usuario")
	}

	if request.EmailVerified {
		if err = userRepository.SetEmailVerified(newUser.ID); err != nil {
			fiber.NewError(fiber.StatusInternalServerError, "No se pudo setear la fecha de actualizacion en verificación del correo del usuario.")
		}
	}

	newUser = userRepository.GetUserById(newUser.ID)

	return ctx.JSON(fiber.Map{
		"usuario": newUser,
	})
}
