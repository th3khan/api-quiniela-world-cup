package admin

import (
	"errors"

	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func CreateUser(name string, email string, roleId uint, password string, active bool) (error, models.User) {
	db := database.Connection()
	userRepository := repositories.NewUserRespository(db)

	return userRepository.CreateUser(
		name,
		email,
		roleId,
		password,
		active,
	)
}

func GetUserByEmail(email string) (error, models.User) {
	db := database.Connection()
	userRepository := repositories.NewUserRespository(db)

	user := userRepository.GetUserByEmail(email)
	if user.ID == 0 {
		return errors.New("Usuario no encontrado."), user
	}

	return nil, user
}

func GetUsers(page int, perPage int) []models.User {
	return []models.User{}
}
