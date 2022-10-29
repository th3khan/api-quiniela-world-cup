package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint   `gorm:"primary_key"`
	RoleId          uint   `gorm:"index;notNull"`
	Role            Role   `gorm:"foreignKey:RoleId"`
	Name            string `gorm:"notNull;size:200"`
	Email           string `gorm:"notNull;size:256"`
	Password        string `gorm:"notNull;size:256"`
	EmailVerified   bool   `gorm:"default:false;notNull"`
	Active          bool   `gorm:"default:false;notNull"`
	Image           string `gorm:"size:256; default:/assets/images/profiles/default.png"`
	EmailVerifiedAt time.Time
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt
}
