package repositories

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type PhaseRepository interface {
	CreatePhase(name string) (error, models.Phase)
	GetPhaseById(id int) (error, models.Phase)
}

type phaseRepository struct {
	db *gorm.DB
}

func NewPhaseRepository(db *gorm.DB) phaseRepository {
	repo := phaseRepository{
		db: db,
	}

	return repo
}

func (repo *phaseRepository) CreatePhase(name string) (error, models.Phase) {
	var phase models.Phase

	phase.Name = name

	result := repo.db.Create(&phase)

	return result.Error, phase
}

func (repo *phaseRepository) GetPhaseById(id uint) (error, models.Phase) {
	var phase models.Phase

	result := repo.db.Where("id = ?", id).Find(&phase)

	return result.Error, phase
}
