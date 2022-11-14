package models

import (
	"time"

	"gorm.io/gorm"
)

const PATH_PROFILE_IMAGES = "/assets/images/profiles/"

type User struct {
	ID              uint    `gorm:"primary_key"`
	RoleId          uint    `gorm:"index;notNull"`
	Role            Role    `gorm:"foreignKey:RoleId"`
	Name            string  `gorm:"notNull;size:200"`
	Email           string  `gorm:"notNull;size:256"`
	Password        string  `gorm:"notNull;size:256"`
	EmailVerified   bool    `gorm:"default:false;"`
	Active          bool    `gorm:"default:false;"`
	Image           string  `gorm:"size:256; default:/assets/images/profiles/default.jpg"`
	Points          float32 `gorm:"default:0"`
	EmailVerifiedAt time.Time
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt
}
