package database

import (
	"Jordan-Tech-Companies/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error reading it, using defaults if applicable")
	}

	database_path := os.Getenv("DATABASE_FILE_PATH")
	//uploads_path := os.Getenv("UPLOADS_FOLDER_PATH")
	fmt.Println(database_path)
	database, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.AutoMigrate(&models.Company{})
	/*company := models.Company{
	      Name : "Google",
	      Location :"Amman",
	      LocationURL: "https://maps.app.goo.gl/was2LXQvWw7scD7QA",
	      Size : "Large - International",
	      Logo : uploads_path + "/Google_logo.png",
	  }
	  database.Create(&company)*/
	DB = database
}

func ReturnAllCompanies() []models.Company {
	companies := []models.Company{}
	result := DB.Find(&companies)
	if result.RowsAffected == 0 {
		fmt.Println("No Rows where found! Something went wrong", result.Error)
		return nil
	}
	return companies
}
