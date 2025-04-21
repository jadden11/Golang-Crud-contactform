package database

import (
	"github.com/username/contact-form-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Ganti sesuai user/password Laragon-mu
	dsn := "root:@tcp(127.0.0.1:3306)/contactsdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto migrate table Contact
	db.AutoMigrate(&models.Contact{})
	DB = db
}
