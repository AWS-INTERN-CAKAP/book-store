package entity

import (
	"time"
)

type Book struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Price	    int		  `gorm:"type:int;not null"`
	ImagePath	string 	  `gorm:"type:varchar(255);not null"`
	Description string 	  `gorm:"type:text;not null"`
	Categories  []Category    `gorm:"many2many:book_category;"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
