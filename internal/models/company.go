package models

type Size int
const (
    Startup    Size = iota   // 0
	Small					 //1
    MidSize                  // 2
    Large             // 3
    RegionallyLarge            // 4
    InternationallyLarge            // 5
)

type Tags int 
const (
	SoftwareDevelopment	Tags = iota // 0
	WebDevelopment					// 1	
	MobileDevelopment				// 2	
	AI								// 3	
	MachineLearning					// 4	
	DeepLearning					// 5	
	DataScience						// 6
	CloudComputing					// 7	
	Blockchain						// 8	
	InternetOfThings				// 9	
	Cybersecurity					// 10	
	Robotics						// 11
	EmbeddedSystems					// 12	
	AugmentedReality				// 13	
	VirtualReality					// 14	
	GameDevelopment					// 15	
	ECommerce						// 16	
	HealthTech						// 17	
	FinTech							// 18	
	EdTech							// 19	
	AdTech							// 20	
	MarTech							// 21	
	PrivateCompany					// 22	
	GovermentalCompany				// 23	
	NonProfitCompany				// 24
	Drones							// 25
	InternshipProgrames				// 26
	DataAnalysis					// 27
	Networking						// 28
	Freelance						// 29
	ComputerVision					// 30
	HR								// 31
	
)


type Company struct {
	ID uint `gorm:"primaryKey"`
	Name string 
	Locations []map[string]string //example: [{city:"Amman", url:"https://maps.app.goo.gl/was2LXQvWw7scD7QA"}] the first location is the base one.
	Size Size
	Logo string //card logo path file
	ContactInfo string
	SocialMedia []map[string]string //example: [{name:"Facebook", url:"https://facebook.com/"}]
	EmployeeCount string
	Tags []Tags //example: [AI, Fintech, E-commerce, Healthtech, Gaming, Edtech, etc.], and this will be used for searching and filtering.
	ProfileImage string //profile page image path
	Description string
}


