package repositories

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type TeamRepository interface {
	GetTeams(page int, perPage int) ([]models.Team, int)
	GetTeam(id uint) models.Team
	CreateTeam(name string, active bool, logo string) (error, models.Team)
	UpdateTeam(id uint, name string, active bool, logo string) (error, models.Team)
	DeleteTeam(id uint) error
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

func (repo *teamRepository) GetTeam(id uint) models.Team {
	var team models.Team
	repo.db.Where("id = ?", id).Find(&team)
	return team
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

func (repo *teamRepository) CreateTeam(name string, active bool, logo string) (error, models.Team) {
	var team models.Team
	team.Name = name
	team.Active = active
	team.Logo = logo
	result := repo.db.Create(&team)
	return result.Error, team
}

func (repo *teamRepository) UpdateTeam(id uint, name string, active bool, logo string) (error, models.Team) {
	var team models.Team
	team.ID = id
	team.Name = name
	team.Active = active
	team.Logo = logo
	result := repo.db.Where("id = ?", id).Updates(&team)
	return result.Error, team
}

func (repo *teamRepository) DeleteTeam(id uint) error {
	result := repo.db.Delete(&models.Team{}, id)
	return result.Error
}
