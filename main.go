package main

import (
	"Go_authentication/config"
	"Go_authentication/models"
	"Go_authentication/router"
	"log"
)

func main() {
	//initialize the database
	config.ConnectDatabase()
	//Migrate  the schema
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}
	
	r := router.SetupRouter()
	r.Run(":8080")
}
