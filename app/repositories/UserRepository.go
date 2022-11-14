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
	DeleteUserById(id int) error
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

	repo.db.Where("email = ?", email).Preload("Role").Find(&user)

	return user
}

func (repo *userRepository) GetUserById(id int) models.User {
	var user models.User

	repo.db.Preload("Role").Where("id = ?", id).Find(&user)

	return user
}

func (repo *userRepository) GetUsers(page int, perPage int) []models.User {
	var users []models.User

	query := repo.db.Model(&models.User{})
	if page > 0 {
		offset := (page - 1) * perPage
		query.Offset(offset).Limit(perPage)
	}

	query.Preload("Role").Find(&users)
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

func (repo *userRepository) UpdateUser(id int, name string, email string, roleId uint, password string, active bool, image string, emailVerified bool, changePasswod bool, changeImage bool, setEmailVerifiedNow bool) error {
	var user models.User

	user.Name = name
	user.Email = email
	user.RoleId = roleId
	user.Active = active

	if changePasswod {
		user.Password = password
	}

	if changeImage {
		user.Image = image
	}

	if setEmailVerifiedNow {
		user.EmailVerified = emailVerified
		user.EmailVerifiedAt = time.Now()
	}

	result := repo.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	return result.Error
}

func (repo *userRepository) DeleteUserById(id int) error {
	result := repo.db.Delete(&models.User{}, id)

	return result.Error
}
