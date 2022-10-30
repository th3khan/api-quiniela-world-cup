package repositories

import (
	"time"

	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name string, email string, roleId uint, password string, active bool) (error, models.User)
	GetUserByEmail(email string) (error, models.User)
	GetUserById(id int) models.User
	GetUsers(page int, perPage int) (error, []models.User)
	GetTotalUsers() (error, int)
	SetEmailVerified(id int) error
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

func (repo *userRepository) CreateUser(name string, email string, roleId uint, password string, active bool, image string, emailVerified bool) (error, models.User) {
	user := models.User{
		RoleId:   roleId,
		Name:     name,
		Email:    email,
		Password: password,
		Active:   active,
	}

	if emailVerified {
		user.EmailVerified = true
		user.EmailVerifiedAt = time.Now()
	}

	if len(image) > 0 {
		user.Image = image
	}

	result := repo.db.Create(&user)

	return result.Error, user
}

func (repo *userRepository) GetUserByEmail(email string) models.User {
	var user models.User

	repo.db.Where("email = ?", email).Find(&user)

	return user
}

func (repo *userRepository) GetUserById(id uint) models.User {
	var user models.User

	repo.db.Where("id = ?", id).Find(&user)

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

func (repo *userRepository) SetEmailVerified(id uint) error {
	var user models.User
	result := repo.db.Model(&user).Where("id = ?", id).Updates(models.User{EmailVerified: true, EmailVerifiedAt: time.Now()})
	return result.Error
}
