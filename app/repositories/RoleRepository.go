package repositories

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(name string) (error, models.Role)
	GetRoleById(id int) (error, models.Role)
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

	role.Name = name

	result := repo.db.Create(&role)

	return result.Error, role
}

func (repo *roleRepository) GetRoleById(id uint) (error, models.Role) {
	var role models.Role

	result := repo.db.Where("id = ?", id).Find(&role)

	return result.Error, role
}
