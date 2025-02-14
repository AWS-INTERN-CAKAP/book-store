package db

import (
	"fmt"

	"github.com/aws-cakap-intern/book-store/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(config *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		`%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return db, err
	}
	return db, nil
}