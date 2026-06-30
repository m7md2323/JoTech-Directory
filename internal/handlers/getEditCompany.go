package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/internal/database"
	"log"
)

func GetEditCompany(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	company, err := database.ReturnCompanyByName(name)
	if err != nil {
		log.Println("Something went wrong inside GetEditCompany handler", err)
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}
	page := pages.GetEditCompany(company)
	page.Render(r.Context(), w)
}
