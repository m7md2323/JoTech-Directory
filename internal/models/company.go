package models

import (
	"time"

)

type Size string
const (
    Startup    				Size = "Startup"
	Small					Size = "Small"
    Medium                  Size = "Medium"
    Large                   Size = "Large"
	Enterprise              Size = "Enterprise"
	MultiNational           Size = "Multi-National"
)

type Type string
const (
    PrivateCompany          Type = "Private Company"
	GovernmentalCompany     Type = "Governmental Company"
	NonProfit               Type = "Non-Profit"
)


type Tags string
const (
	SoftwareDevelopment		Tags = "Software Development"
	WebDevelopment			Tags = "Web Development"	
	MobileDevelopment		Tags = "Mobile Development"	
	AI						Tags = "AI"	
	MachineLearning			Tags = "Machine Learning"	
	DeepLearning			Tags = "Deep Learning"	
	DataScience				Tags = "Data Science"       
	CloudComputing			Tags = "Cloud Computing"     
	Blockchain				Tags = "Blockchain"         
	InternetOfThings		Tags = "Internet Of Things"    
	Cybersecurity			Tags = "Cyber Security"      
	Robotics				Tags = "Robotics"           
	EmbeddedSystems			Tags = "Embedded Systems"    
	AugmentedReality		Tags = "Augmented Reality"  
	VirtualReality			Tags = "Virtual Reality"    
	GameDevelopment			Tags = "Game Development"  
	ECommerce				Tags = "ECommerce"          
	HealthTech				Tags = "Health Tech"          
	FinTech					Tags = "FinTech"           
	EdTech					Tags = "EdTech"            
	AdTech					Tags = "AdTech"           
	MarTech					Tags = "MarTech"          
	Drones					Tags = "Drones"             
	InternshipPrograms		Tags = "Internship Programs" 
	DataAnalysis			Tags = "Data Analysis"       
	Networking				Tags = "Networking"         
	Freelance				Tags = "Freelance"          
	ComputerVision			Tags = "Computer Vision"     
	HumanResources			Tags = "Human Resources"
)

type Location struct {
    City string
    URL  string
}

type Link struct {
    Platform string
    URL      string
}


type Company struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Locations []Location `gorm:"serializer:json"` //example: [{city:"Amman", url:"https://maps.app.goo.gl/was2LXQvWw7scD7QA"}] the first location is the base one.
	Size Size
	Logo string //card logo path file
	ContactInfo string
	Links []Link `gorm:"serializer:json"` //example: [{name:"Facebook", url:"https://facebook.com/"}]
	EmployeeCount string
	Tags []Tags `gorm:"serializer:json"` //example: [AI, Fintech, E-commerce, Healthtech, Gaming, Edtech, etc.], and this will be used for searching and filtering.
	ProfileImage string //profile page image path
	Description string `gorm:"type:text"`
	Type Type
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}


