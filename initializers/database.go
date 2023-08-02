package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// In order to make DB accessible from other packages
var DB *gorm.DB

// to set up the Gorm database connection
func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db
	log.Println("Connected to the database.")
}
