package handlers

import (
	"github.com/m7md2323/JoTech-Directory/internal/database"
	"github.com/m7md2323/JoTech-Directory/internal/models"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func PostAddCompany(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	

	/*
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

	newCompany := models.Company{
		Name:   r.FormValue("name"),
		Locations:     locations,
		Size:          models.Size(r.FormValue("size")),
		Type:          models.Type(r.FormValue("type")),
		Links:         links,
		ContactInfo:   r.FormValue("contactInfo"),
		EmployeeCount: r.FormValue("employeeCount"),
		Tags:          tags,
		Description:   r.FormValue("description"),
		Logo:          newLogoFileName,
	}

	database.DB.Create(&newCompany)

	page := pages.PostAddCompany()
	page.Render(r.Context(), w)
}

func saveFile(file multipart.File, path string) {
	

	uploadsPath := os.Getenv("UPLOADS_FOLDER_PATH")

	newFile, err := os.OpenFile(uploadsPath+"/"+path,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println("os Something went wrong", err)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		log.Println("io Something went wrong", err)
		return
	}
}
