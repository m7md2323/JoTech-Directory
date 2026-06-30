package handlers

import (
	"github.com/m7md2323/JoTech-Directory/internal/database"
	"github.com/m7md2323/JoTech-Directory/internal/models"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
	"log"
	"net/http"
	"path/filepath"
)

func PostAddEvent(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	
	/*
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
	
	*/

	locations := []models.Location{}
	links := []models.Link{}
	tags := []models.Tag{}

	tagsString := r.Form["tags"]

	for i := range tagsString {
		tag:=models.Tag{
			Name: tagsString[i],
		}
		tags = append(tags, tag)
	}

	baseCity := r.FormValue("city")
	baseCiryURL := r.FormValue("locationUrl")

	locations = append(locations, models.Location{
		City: baseCity,
		URL:  baseCiryURL,
	})

	otherCities := r.Form["other_cities"]
	otherCitiesURLs:=r.Form["other_locationUrls"]

	for i := range otherCities {
		locations = append(locations, models.Location{
			City: otherCities[i],
			URL:  otherCitiesURLs[i],
		})
	}

	website := r.FormValue("website")
	linkedin := r.FormValue("linkedin")
	twitter := r.FormValue("twitter")
	facebook := r.FormValue("facebook")

	links = append(links, models.Link{
		Platform: "website",
		URL:      website,
	})
	links = append(links, models.Link{
		Platform: "linkedin",
		URL:      linkedin,
	})
	links = append(links, models.Link{
		Platform: "twitter",
		URL:      twitter,
	})
	links = append(links, models.Link{
		Platform: "facebook",
		URL:      facebook,
	})

	logo, logoHeader, err1 := r.FormFile("logo")
	if err1 != nil {
		log.Println(err1)
		return
	}
	defer logo.Close()

	profileImage, ProfileHeader, err2 := r.FormFile("profileImage")
	if err2 != nil {
		log.Println(err2)
	}
	defer profileImage.Close()

	newLogoFileName :=  r.FormValue("name") + "Logo" + filepath.Ext(logoHeader.Filename)
	newProfileImageFileName := r.FormValue("name") + "ProfileImage" + filepath.Ext(ProfileHeader.Filename)

	saveFile(logo, newLogoFileName)
	saveFile(profileImage, newProfileImageFileName)

	newCompany := models.Event{
		Name:   r.FormValue("name"),
		Locations:     locations,
		Type:          models.Type(r.FormValue("type")),
		Links:         links,
		ContactInfo:   r.FormValue("contactInfo"),
		Tags:          tags,
		Description:   r.FormValue("description"),
		Logo:          newLogoFileName,
		ProfileImage:  newProfileImageFileName,
	}

	database.DB.Create(&newCompany)

	page := pages.MessageAddEvent()
	renderErr := page.Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ", renderErr)
	}
}
