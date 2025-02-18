package entity

import "time"

type BookCategory struct {
	BookID     uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}