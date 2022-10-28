package user

import (
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
