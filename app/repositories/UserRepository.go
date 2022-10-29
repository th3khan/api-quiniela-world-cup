package repositories

import (
	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name string, email string, roleId uint, password string, active bool) (error, models.User)
	GetUserByEmail(email string) (error, models.User)
	GetUsers(page int, perPage int) (error, []models.User)
	GetTotalUsers() (error, int)
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

func (repo *userRepository) GetUserByEmail(email string) models.User {
	var user models.User

	repo.db.Where("email = ?", email).Find(&user)

	return user
}

func (repo *userRepository) GetUsers(page int, perPage int) []models.User {
	var users []models.User

	query := repo.db
	if page > 0 {
		query.Offset(page).Limit(perPage)
	}

	query.Find(&users)
	return users
}

func (repo *userRepository) GetTotalUsers() int {
	var total int
	var users []models.User

	repo.db.Find(&users)

	total = len(users)

	return total
}
