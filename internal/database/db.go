package database

import (
	"github.com/m7md2323/JoTech-Directory/internal/models"
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
	database, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.Exec("PRAGMA foreign_keys = ON;")

	migrateError:= database.AutoMigrate(&models.Company{})
	if migrateError != nil {
		log.Println("Something went wrong while AutoMigrating Company model",migrateError)
	}
	migrateError = database.AutoMigrate(&models.Tag{})
	if migrateError != nil {
		log.Println("Something went wrong while AutoMigrating Tag model",migrateError)
	}
	//database.AutoMigrate(&models.Event{})
	migrateError = database.AutoMigrate(&models.Location{})
	if migrateError != nil {
		log.Println("Something went wrong while AutoMigrating Location model",migrateError)
	}
	DB = database
}

func ReturnAllCompanies() ([]models.Company,error) {
	companies := []models.Company{}
	err := DB.Where("status = ?", "published").Preload("Locations").Preload("Tags").Find(&companies).Error
	return companies,err
}

func ReturnAllEvents() ([]models.Event,error) {
	events := []models.Event{}
	err := DB.Model(&models.Event{}).Preload("Locations").Preload("Tags").Find(&events).Error
	return events,err
}
func ReturnCompaniesByQuery(paramas  models.FilterParams) ([]models.Company, error) {

	var companies []models.Company
	query := DB.Model(&models.Company{}).Preload("Locations").Preload("Tags")
	if paramas.SearchTerm != "" {
		query = query.Where("name LIKE ?","%" + paramas.SearchTerm + "%")		
	}

	if len(paramas.Sizes) >0 && paramas.Sizes[0] != "" {
		query=query.Where("size IN ?",paramas.Sizes)
	}

	if len(paramas.Types) >0 && paramas.Types[0] != ""{
		query=query.Where("type IN ?",paramas.Types)
	}

	if len(paramas.Cities) > 0 && paramas.Cities[0] != ""{
		query=query.Where("id IN (SELECT company_id FROM locations where city IN ?)",paramas.Cities)
	}

	if len(paramas.Tags) > 0 && paramas.Tags[0] != ""{
		query=query.Where("id IN (SELECT company_id FROM tags where name IN ?)",paramas.Tags)
	}

	if paramas.Status == "draft" {
		query = query.Where("status = ?", "draft")
	} else {
		query = query.Where("status = ?", "published")
	}

	
	err := query.Debug().Find(&companies).Error
	return companies,err
}

func ReturnCompanyByName(name string) (models.Company,error) {
	var company models.Company
	err := DB.Model(&models.Company{}).Preload("Locations").Preload("Tags").Where("name = ?",name).First(&company).Error
	return company,err
}

