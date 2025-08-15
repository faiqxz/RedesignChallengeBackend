package database

import (
	"log"

	"github.com/glebarez/sqlite" // Import the pure-Go driver
	"gorm.io/gorm"
	"redesign/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("fasilkom.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!\n", err)
	}

	log.Println("Running Migrations")
	DB.AutoMigrate(

		&models.News{},
		&models.ResearchTeam{},
		&models.DownloadableFile{},
		&models.Certification{},
		&models.Gallery{},
	)

	log.Println("Database connection successfully opened")
}
