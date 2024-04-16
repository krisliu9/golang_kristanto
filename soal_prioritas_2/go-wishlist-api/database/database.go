package database

import (
	"go-wishlist-api/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect to the database
func InitDatabase() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&repository.Wishlist{})
	DB.AutoMigrate(&repository.User{})
}
