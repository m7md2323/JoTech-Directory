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
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Size Size //startup, small, medium, large, enterprise, multi-national
	Logo string //card logo path file
	ContactInfo string 
	EmployeeCount string //example: "1-10", "11-50", "51-200", "201-500"
	ProfileImage string //profile page image path
	Description string `gorm:"type:text"`
	Type Type //Private Company, Governmental Company, Non-Profit

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Locations []Location `gorm:"foreignKey:CompanyID"` //example: [{city:"Amman", url:"example.com"}] the first location is the base one.
	Tags []Tag `gorm:"foreignKey:CompanyID"` //example: [AI, Fintech, E-commerce, Healthtech, Gaming, Edtech, etc.], and this will be used for searching and filtering.
	Links []Link `gorm:"serializer:json"` //example: [{name:"Facebook", url:"https://facebook.com/"}]

}



