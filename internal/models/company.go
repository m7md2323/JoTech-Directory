package models

import (
	//"time"
	"gorm.io/gorm"

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



type Link struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`

}

type FilterParams struct {
	SearchTerm string   // Text search input
	Sizes      []string // From checkboxes: ["Small", "Medium"]
	Types      []string // From checkboxes: ["Private Company"]
	Cities     []string // From checkboxes: ["Amman", "Irbid"]
	Tags       []string // From checkboxes: ["AI", "Web Development"]
	Status     string   // From toggle switch: "draft"
}

/*type Tags string
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

*/

type Company struct {
    gorm.Model
    Name              string     `gorm:"unique;not null"`
    Size 			  Size
    Logo              string
    ContactInfo       string     `gorm:"column:contact_info"`
    EmployeeCount     string     `gorm:"column:employee_count"`
    Description       string
    Type 			  Type 
    Links []Link 				 `gorm:"serializer:json"`
    InternshipProgram string     `gorm:"column:internship_program"`
    Status            string     `gorm:"default:draft"`
    Locations         []Location `gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE;"`
    Tags              []Tag      `gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE;"`
}
