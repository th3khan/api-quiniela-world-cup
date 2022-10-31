package repositories

import (
	"strings"

	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type IGroupRepository interface {
	GetGroups(page int, perPage int) ([]models.Group, int)
	CreateGroup(name string, active bool) (error, models.Group)
	UpdateGroup(id int, name string, active bool) (error, models.Group)
	GetGroupById(id int) models.Group
	GetGroupByName(name string) models.Group
	DeleteGroupById(id int) error
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) groupRepository {
	repo := groupRepository{
		db: db,
	}
	return repo
}

func (repo *groupRepository) GetGroups(page int, perPage int) ([]models.Group, int) {
	var groups []models.Group

	query := repo.db.Model(&models.Group{})
	if page > 0 {
		offset := (page - 1) * perPage
		query.Offset(offset).Limit(perPage)
	}
	query.Find(&groups)
	return groups, len(repo.GetAllGroups())
}

func (repo *groupRepository) GetAllGroups() []models.Group {
	var groups []models.Group
	repo.db.Find(&groups)
	return groups
}

func (repo *groupRepository) CreateGroup(name string, active bool) (error, models.Group) {
	var group models.Group
	group.Name = strings.ToUpper(name)
	group.Active = active
	result := repo.db.Create(&group)
	return result.Error, group
}

func (repo *groupRepository) UpdateGroup(id int, name string, active bool) (error, models.Group) {
	var group models.Group
	group.Name = strings.ToUpper(name)
	group.Active = active
	result := repo.db.Where("id = ?", id).Updates(&group)
	return result.Error, group
}

func (repo *groupRepository) GetGroupById(id int) models.Group {
	var group models.Group
	repo.db.Where("id = ?", id).Find(&group)
	return group
}

func (repo *groupRepository) GetGroupByName(name string, excludeId int) models.Group {
	var group models.Group
	query := repo.db.Where("name = ?", strings.ToUpper(name))

	if excludeId > 0 {
		query.Where("id <> ?", excludeId)
	}

	query.Find(&group)
	return group
}

func (repo *groupRepository) DeleteGroupById(id int) error {
	result := repo.db.Delete(&models.Group{}, id)
	return result.Error
}
