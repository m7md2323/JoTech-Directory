package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/models"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/web/templates/partials"
	"fmt"
	"net/http"
)

func Companies(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	companies,err:=database.ReturnAllCompanies()
	if err !=nil {
		fmt.Println("Something went wrong in Companies handler",err)
	}
	page := pages.Companies(companies)
	page.Render(r.Context(), w)

}

func SearchAndFilterCompanies(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	Paramas := models.FilterParams{
		SearchTerm: r.FormValue("search"),
		Sizes:      r.Form["size"],
		Types:      r.Form["type"],
		Cities:     r.Form["city"],
		Tags:       r.Form["tag"],
	}	

	companiesRes,err:=database.ReturnCompaniesByQuery(Paramas)

	if(err!=nil){
		fmt.Println("Something went wrong in Searching and Filtering",err)
		return
	}

	w.WriteHeader(http.StatusOK)
	partial := partials.CompaniesList(companiesRes)
	partial.Render(r.Context(), w)

}
