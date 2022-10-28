package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint   `gorm:"primary_key"`
	RoleId          uint   `gorm:"index;notNull"`
	Role            Role   `gorm:"foreignKey:RoleId"`
	Name            string `gorm:"notNull"`
	Email           string `gorm:"notNull"`
	Password        string `gorm:"notNull"`
	EmailVerified   bool   `gorm:"default:false;notNull"`
	Active          bool   `gorm:"default:false;notNull"`
	EmailVerifiedAt time.Time
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt
}
