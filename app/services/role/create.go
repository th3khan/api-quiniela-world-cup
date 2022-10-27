package role

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func CreateRole(name string) (error, models.Role) {
	db := database.Connection()
	roleRepository := repositories.NewRoleRepository(db)

	err, role := roleRepository.CreateRole(name)

	return err, role
}
