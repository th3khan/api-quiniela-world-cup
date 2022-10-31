package models

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"notNull;size:200"`
	Active    bool      `gorm:"notNull"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
