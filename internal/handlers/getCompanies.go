package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/models"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/web/templates/partials"
	"fmt"
	"net/http"
	"sort"
)

func Companies(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	companies,err:=database.ReturnAllCompanies()
	if err !=nil {
		fmt.Println("Something went wrong in Companies handler",err)
	}

	sort.Slice(companies, func(i, j int) bool {
		return companies[i].Name < companies[j].Name
	})

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
		Status:     r.FormValue("status"),
	}	

	companiesRes,err:=database.ReturnCompaniesByQuery(Paramas)

	if(err!=nil){
		fmt.Println("Something went wrong in Searching and Filtering",err)
		return
	}

	sort.Slice(companiesRes, func(i, j int) bool {
		return companiesRes[i].Name < companiesRes[j].Name
	})
	w.WriteHeader(http.StatusOK)
	partial := partials.CompaniesList(companiesRes)
	partial.Render(r.Context(), w)

}
