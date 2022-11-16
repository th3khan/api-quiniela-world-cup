package models

import (
	"time"

	"gorm.io/gorm"
)

const STATUS_PAYMENT_PENDING = 0
const STATUS_PAYMENT_APPROVED = 1
const STATUS_PAYMENT_CANCELED = 2
const PATH_FOLDER_PAYMENTS = "/assets/images/payments/"

type Payment struct {
	ID        uint    `gorm:"primaryKey"`
	UserId    uint    `gorm:"index"`
	User      User    `gorm:"foreignKey:UserId"`
	Amount    float32 `gorm:"notNull"`
	Status    int32   `gorm:"default:0"`
	Date      time.Time
	Image     string    `gorm:"size:256"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
