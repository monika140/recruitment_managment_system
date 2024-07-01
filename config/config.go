package config

import (
	"Go_authentication/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=monika11 dbname=auth port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	// AutoMigrate to create tables if they do not exist
	database.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{})
	DB = database
}
