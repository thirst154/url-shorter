package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("POSTGRES_CONNECTION")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}

	// Independent tables first, then tables with foreign key dependencies.
	err = DB.AutoMigrate(
		&URL{},
	)
	if err != nil {
		panic("Failed to auto migrate DB")
	}
}
