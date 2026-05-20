package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/models"
)

func GetCompanyProfile(w http.ResponseWriter, r *http.Request) {

	name :=r.PathValue("name")
	company := models.Company{}
	database.DB.Where("name = ?",name).First(&company)

	page:= pages.CompanyProfile(company)
	page.Render(r.Context(),w)
	w.WriteHeader(http.StatusOK)
}
