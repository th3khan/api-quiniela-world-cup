package repositories

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name string, email string, roleId uint, password string, active bool) (error, models.User)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRespository(db *gorm.DB) userRepository {
	repo := userRepository{
		db: db,
	}

	return repo
}

func (repo *userRepository) CreateUser(name string, email string, roleId uint, password string, active bool) (error, models.User) {
	user := models.User{
		RoleId:   roleId,
		Name:     name,
		Email:    email,
		Password: password,
		Active:   active,
	}

	result := repo.db.Create(&user)

	return result.Error, user
}
