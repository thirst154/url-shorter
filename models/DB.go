package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	if err = DB.AutoMigrate(&URL{}); err != nil {
		log.Fatal("Failed to migrate DB: ", err)
	}
}
