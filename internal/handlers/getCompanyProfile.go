package handlers

import (
	"github.com/m7md2323/JoTech-Directory/internal/database"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
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
