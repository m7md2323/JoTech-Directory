package models

import (
	"gorm.io/gorm"

)

type Event struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Logo string //card logo path file
	ContactInfo string 
	ProfileImage string //profile page image path
	Description string `gorm:"type:text"`
	Type Type //Private Company, Governmental Company, Non-Profit

	Locations []Location `gorm:"foreignKey:CompanyID"` //example: [{city:"Amman", url:"example.com"}] the first location is the base one.
	Tags []Tag `gorm:"foreignKey:CompanyID"` //example: [AI, Fintech, E-commerce, Healthtech, Gaming, Edtech, etc.], and this will be used for searching and filtering.
	Links []Link `gorm:"serializer:json"` //example: [{name:"Facebook", url:"https://facebook.com/"}]

}


