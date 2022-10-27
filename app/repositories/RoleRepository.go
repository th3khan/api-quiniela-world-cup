package repositories

import (
	"errors"

	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(name string) (error, models.Role)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) roleRepository {
	repo := roleRepository{
		db: db,
	}

	return repo
}

func (repo *roleRepository) CreateRole(name string) (error, models.Role) {
	var role models.Role

	if len(name) <= 0 {
		return errors.New("Name is required"), role
	}

	role.Name = name

	result := repo.db.Create(&role)

	return result.Error, role
}
