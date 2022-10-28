package seeders

import (
	"log"

	"github.com/th3khan/api-quiniela-world-cup/app/helpers"
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/services/user"
)

func (s *Seed) CreateUsers() error {
	passwordDefault, err := helpers.HashingPassword("123456")
	if err != nil {
		return err
	}
	users := []models.User{
		models.User{
			ID:            1,
			RoleId:        models.ROLE_SUPER_ADMIN,
			Name:          "Super admin",
			Email:         "superadmin@mail.com",
			Password:      passwordDefault,
			Active:        true,
			EmailVerified: true,
		},
	}

	for _, u := range users {
		if err, _ := user.CreateUser(
			u.Name,
			u.Email,
			u.RoleId,
			u.Password,
			u.Active,
		); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
