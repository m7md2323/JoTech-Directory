package models

import (
	"time"

)

type Event struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Logo string //card logo path file
	ContactInfo string 
	ProfileImage string //profile page image path
	Description string `gorm:"type:text"`
	Type Type //Private Company, Governmental Company, Non-Profit

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Locations []Location `gorm:"foreignKey:CompanyID"` //example: [{city:"Amman", url:"example.com"}] the first location is the base one.
	Tags []Tag `gorm:"foreignKey:CompanyID"` //example: [AI, Fintech, E-commerce, Healthtech, Gaming, Edtech, etc.], and this will be used for searching and filtering.
	Links []Link `gorm:"serializer:json"` //example: [{name:"Facebook", url:"https://facebook.com/"}]

}


