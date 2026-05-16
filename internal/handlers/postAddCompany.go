package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/models"
	"Jordan-Tech-Companies/web/templates/pages"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

func PostAddCompany(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	

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
	tags := []models.Tags{}

	locationsString := r.Form["locations"]
	locationUrlString := r.Form["locationUrl"]
	tagsString := r.Form["tags"]

	for i := range locationsString {
		locations = append(locations, models.Location{
			City: locationsString[i],
			URL:  locationUrlString[i],
		})
	}

	for i := range tagsString {
		tags = append(tags, models.Tags(tagsString[i]))
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
	}
	defer logo.Close()

	profileImage, ProfileHeader, err2 := r.FormFile("profileImage")
	if err2 != nil {
		log.Println(err2)
	}
	defer profileImage.Close()

	newLogoFileName :=  r.FormValue("name") + "Logo" + filepath.Ext(logoHeader.Filename)
	newProfileImageFileName := r.FormValue("name") + "ProfileImage" + filepath.Ext(ProfileHeader.Filename)

	fmt.Println(newLogoFileName)

	saveFile(logo, newLogoFileName)
	saveFile(profileImage, newProfileImageFileName)

	newCompany := models.Company{
		Name:          r.FormValue("name"),
		Locations:     locations,
		Size:          models.Size(r.FormValue("size")),
		Type:          models.Type(r.FormValue("type")),
		Links:         links,
		ContactInfo:   r.FormValue("contactInfo"),
		EmployeeCount: r.FormValue("employeeCount"),
		Tags:          tags,
		Description:   r.FormValue("description"),
		Logo:          newLogoFileName,
		ProfileImage:  newProfileImageFileName,
	}

	database.DB.Create(&newCompany)

	page := pages.PostAddCompany()
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)
}

func saveFile(file multipart.File, path string) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error reading it, using defaults if applicable")
	}

	uploadsPath := os.Getenv("UPLOADS_FOLDER_PATH")

	newFile, err := os.OpenFile(uploadsPath+"/"+path,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println("os Something went wrong", err)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Println("io Something went wrong", err)
		return
	}
}
