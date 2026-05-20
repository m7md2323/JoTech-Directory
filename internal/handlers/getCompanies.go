package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/web/templates/partials"
	"fmt"
	"net/http"
)

func Companies(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		page := pages.Companies(database.ReturnAllCompanies())
		page.Render(r.Context(), w)
	
}

func SearchCompanies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	companyName := r.FormValue("search")
	fmt.Println("Received search query:", companyName)
	w.WriteHeader(http.StatusOK)
	
	if companyName == "" {
		partial := partials.CompaniesList(database.ReturnAllCompanies())
		partial.Render(r.Context(), w)

	}else{
		partial:= partials.CompaniesList(database.ReturnCompaniesByName(companyName))
		partial.Render(r.Context(), w)
	}
}


