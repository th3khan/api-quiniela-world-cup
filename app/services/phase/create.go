package phase

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"github.com/th3khan/api-quiniela-world-cup/app/repositories"
	"github.com/th3khan/api-quiniela-world-cup/platform/database"
)

func CreatePhase(name string) (models.Phase, error) {
	db := database.Connection()
	repo := repositories.NewPhaseRepository(db)

	err, phase := repo.CreatePhase(name)

	return phase, err
}
