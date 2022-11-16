package migrations

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func Migrate() {
	database.Connection().AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Group{},
		&models.Team{},
		&models.Phase{},
		&models.Payment{},
	)
}
