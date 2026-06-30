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

	//"gorm.io/gorm"
)

func PostEditCompany(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return
	}

	companyName := r.FormValue("name")
	company, err := database.ReturnCompanyByName(companyName)

	if err != nil {
		log.Println("Error retrieving company by name: ", err)
		http.Error(w, "Company not found", http.StatusNotFound)
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
		tag := models.Tag{
			Name: tagsString[i],
		}

		tags = append(tags, tag)
	}

	baseCity := r.FormValue("city")
	baseCityURL := r.FormValue("locationUrl")

	otherCities := r.Form["other_cities"]
	otherCitiesURLs := r.Form["other_locationUrls"]

	locations = append(locations, models.Location{
		City: baseCity,
		URL:  baseCityURL,
	})

	for i := range otherCities {

		locations = append(locations, models.Location{
			City: otherCities[i],
			URL:  otherCitiesURLs[i],
		})

	}

	website := r.FormValue("platform")
	websiteURL := r.FormValue("linkUrl")

	otherPlatforms := r.Form["other_platforms"]
	otherPlatformsURLs := r.Form["other_linkUrls"]

	links = append(links, models.Link{
		Platform: website,
		URL:      websiteURL,
	})

	for i := range otherPlatforms {

		links = append(links, models.Link{
			Platform: otherPlatforms[i],
			URL:      otherPlatformsURLs[i],
		})

	}

	logo, logoHeader, logoErr := r.FormFile("logo")
	if logoErr != nil {
		log.Println("Logo error", logoErr)
	} else {
		defer logo.Close()
	}

	var newLogoFileName string

	if logoErr == nil {
		newLogoFileName = r.FormValue("name") + "Logo" + filepath.Ext(logoHeader.Filename)
		saveFile(logo, newLogoFileName)
	} else {
		newLogoFileName = company.Logo
	}

	

	//First, update the main fields of the company
	updateErr:=database.DB.Model(&company).Updates(models.Company{

        Size:          models.Size(r.FormValue("size")),
        Links:         links,
        ContactInfo:   r.FormValue("contactInfo"),
        EmployeeCount: r.FormValue("employeeCount"),
        Description:   r.FormValue("description"),
        Logo:          newLogoFileName,
        Status:        "published",
    }).Error

    if updateErr !=nil {
        log.Println("Something went wrong in updating "+companyName,updateErr)
    }

    updateErr = database.DB.Model(&company).Association("Locations").Replace(locations)

    if updateErr != nil {
        log.Println("Something went wrong in updating " + companyName + " in Associating with locations: ",updateErr)
    }

    updateErr = database.DB.Model(&company).Association("Tags").Replace(tags)

    if updateErr != nil {
        log.Println("Something went wrong in updating " + companyName + " in Associating with tags: ",updateErr)
    } 


	/*updateErr := database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Update the main Company fields 
		// (Using a map is safer here to ensure empty strings from the form are actually updated)
		if err := tx.Model(&company).Updates(map[string]interface{}{
			"size":           models.Size(r.FormValue("size")),
			"links":          links,
			"contact_info":   r.FormValue("contactInfo"),
			"employee_count": r.FormValue("employeeCount"),
			"description":    r.FormValue("description"),
			"logo":           newLogoFileName,
			"profile_image":  newProfileImageFileName,
			"status":         "published",
		}).Error; err != nil {
			fmt.Println("Phase one")
			return err // Rollback
		}

		// 2. EXPLICIT WIPE: Permanently delete the old associations for this company
		// Using Unscoped() prevents soft-delete bloat in your database over time
		if err := tx.Unscoped().Where("company_id = ?", company.ID).Delete(&models.Location{}).Error; err != nil {
			fmt.Println("Phase two")
			return err
		}
		if err := tx.Unscoped().Where("company_id = ?", company.ID).Delete(&models.Tag{}).Error; err != nil {
			fmt.Println("Phase three")
			return err
		}

		// 3. DATA SCRUB: Ensure the incoming data is strictly formatted for a fresh INSERT
		for i := range locations {
			locations[i].ID = 0                 // Strip any rogue IDs from the form so SQLite doesn't panic
			locations[i].CompanyID = company.ID // Explicitly satisfy the Foreign Key
		}
		for i := range tags {
			tags[i].ID = 0
			tags[i].CompanyID = company.ID
		}

		// 4. EXPLICIT INSERT: Create the new clean records
		if len(locations) > 0 {
			if err := tx.Create(&locations).Error; err != nil {
				return err
			}
		}
		
		if len(tags) > 0 {
			if err := tx.Create(&tags).Error; err != nil {
				return err
			}
		}

		return nil // Commit the transaction
	})*/
	

	if updateErr != nil {
		log.Printf("Failed to fully update company %s: %v", companyName, updateErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	page := pages.PostEditCompany()
	page.Render(r.Context(), w)
}

func saveFilee(file multipart.File, path string) {
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
