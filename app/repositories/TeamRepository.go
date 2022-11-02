package repositories

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type TeamRepository interface {
	GetTeams(page int, perPage int) ([]models.Team, int)
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) teamRepository {
	repo := teamRepository{
		db: db,
	}
	return repo
}

func (repo *teamRepository) GetTeams(page int, perPage int) ([]models.Team, int) {
	var teams []models.Team

	query := repo.db.Model(&models.Team{})
	if page > 0 {
		offset := (page - 1) * perPage
		query.Offset(offset).Limit(perPage)
	}
	query.Find(&teams)
	return teams, len(repo.GetAllTeams())
}

func (repo *teamRepository) GetAllTeams() []models.Team {
	var teams []models.Team
	repo.db.Find(&teams)
	return teams
}
