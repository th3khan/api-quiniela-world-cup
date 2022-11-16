package repositories

import (
	"time"

	"github.com/th3khan/api-quiniela-world-cup/app/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(date time.Time, amount float32, image string, userId uint) (error, models.Payment)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) paymentRepository {
	repo := paymentRepository{
		db: db,
	}
	return repo
}

func (repo *paymentRepository) CreatePayment(date time.Time, amount float32, image string, userId uint) (error, models.Payment) {
	payment := &models.Payment{
		UserId: userId,
		Amount: amount,
		Date:   date,
		Status: models.STATUS_PAYMENT_PENDING,
		Image:  image,
	}

	result := repo.db.Create(payment)
	return result.Error, *payment
}
