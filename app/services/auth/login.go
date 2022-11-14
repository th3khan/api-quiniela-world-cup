package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func Login(email string, password string) (error, string, *models.User) {
	db := database.Connection()
	repo := repositories.NewUserRespository(db)

	user := repo.GetUserByEmail(email)

	if user.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Credenciales invalidas."), "", &models.User{}
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Credenciales invalidas"), "", &models.User{}
	}

	token, err := createToken(&user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error()), "", &models.User{}
	}

	return nil, token, &user
}

func createToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"name":       user.Name,
		"email":      user.Email,
		"roleId":     user.RoleId,
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	}

	return helpers.GenerateToken(claims)

}
