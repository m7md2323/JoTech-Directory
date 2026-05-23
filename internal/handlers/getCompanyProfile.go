package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/web/templates/pages"
	"fmt"
	"net/http"
)

func GetCompanyProfile(w http.ResponseWriter, r *http.Request) {

	name := r.PathValue("name")
	company, err := database.ReturnCompanyByName(name)
	if err != nil {
		fmt.Println("Something went wrong inside GetCompanyProfile handler", err)
	}
	w.WriteHeader(http.StatusOK)
	page := pages.CompanyProfile(company)
	page.Render(r.Context(), w)
}
