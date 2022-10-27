package models

import "time"

type Role struct {
	ID        uint      `gorm:"promaryKey"`
	Name      string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
