package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/web/templates/pages"
	"log"
	"net/http"
)

func GetCompanyProfile(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	company, err := database.ReturnCompanyByName(name)
	if err != nil {
		log.Println("Something went wrong inside GetCompanyProfile handler", err)
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}
	page := pages.CompanyProfile(company)
	page.Render(r.Context(), w)
}
