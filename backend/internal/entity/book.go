package entity

import (
	"time"
)

type Book struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Title       string     `gorm:"type:varchar(255);not null"`
	Price       int        `gorm:"type:int;not null"`
	ImagePath   string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text;not null"`
	Categories  []Category `gorm:"many2many:book_categories;"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}
