package models

import (
	"time"

	"gorm.io/gorm"
)

const PATH_FOLDER_LOGO_TEAMS = "/assets/images/teams/"

type Team struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"notNull;size:200"`
	Active    bool      `gorm:"notNull;default:false"`
	Logo      string    `gorm:"notNull;size:255;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
