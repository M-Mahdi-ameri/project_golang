package database

import (
	"log"

	"github.com/mahdi/library-api/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("faild to connect to database", err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatal("faild to mirage database: ", err)
	}

	DB = database
}
