package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/internal/database"
	"fmt"
)

func GetEditCompany(w http.ResponseWriter, r *http.Request) {


	name := r.PathValue("name")
	company, err := database.ReturnCompanyByName(name)
	if err != nil {
		fmt.Println("Something went wrong inside GetEditCompany handler", err)
	}

	fmt.Println("DEBUG company.ID:", company.ID)
	fmt.Println("DEBUG company.Name:", company.Name)
	fmt.Println("DEBUG company.Locations count:", len(company.Locations))
	fmt.Println("DEBUG company.Locations:", company.Locations)

	w.WriteHeader(http.StatusOK)
	page := pages.GetEditCompany(company)
	page.Render(r.Context(), w)

}
