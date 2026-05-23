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
	fmt.Println(database_path)
	database, err := gorm.Open(sqlite.Open(database_path), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.AutoMigrate(&models.Company{})
	database.AutoMigrate(&models.Tag{})
	database.AutoMigrate(&models.Event{})
	database.AutoMigrate(&models.Location{})
	DB = database
}

func ReturnAllCompanies() ([]models.Company,error) {
	companies := []models.Company{}
	err := DB.Preload("Locations").Preload("Tags").Find(&companies).Error
	return companies,err
}

func ReturnAllEvents() ([]models.Event,error) {
	events := []models.Event{}
	err := DB.Preload("Locations").Preload("Tags").Find(&events).Error
	return events,err
}
func ReturnCompaniesByQuery(paramas  models.FilterParams) ([]models.Company, error) {

	var companies []models.Company
	query := DB.Model(&models.Company{}).Preload("Locations").Preload("Tags")
	fmt.Println("Hello1 ",paramas)
	if paramas.SearchTerm != "" {
		query = query.Where("name LIKE ?","%" + paramas.SearchTerm + "%")		
		fmt.Println(query.RowsAffected)
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

	
	err := query.Debug().Find(&companies).Error
	fmt.Println(companies)	
	return companies,err
}

func ReturnCompanyByName(name string) (models.Company,error) {
	var company models.Company
	err := DB.Preload("Locations").Preload("Tags").First(&company).Where("name = ?",name).Error
	return company,err
}