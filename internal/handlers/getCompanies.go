package handlers

import (
	"github.com/m7md2323/JoTech-Directory/internal/database"
	"github.com/m7md2323/JoTech-Directory/internal/models"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
	"github.com/m7md2323/JoTech-Directory/web/templates/partials"
	"log"
	"net/http"
	"sort"
)

func Companies(w http.ResponseWriter, r *http.Request) {
	companies, err := database.ReturnAllCompanies()
	if err != nil {
		log.Println("Something went wrong in Companies handler", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sort.Slice(companies, func(i, j int) bool {
		return companies[i].Name < companies[j].Name
	})

	page := pages.Companies(companies)
	renderErr:=page.Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ",renderErr)
	}

}

func SearchAndFilterCompanies(w http.ResponseWriter, r *http.Request) {
	parseErr:=r.ParseForm()
	
	if parseErr !=nil {
		log.Println("Something went wrong while parsing in SearchAndFilter ",parseErr)
	}

	Paramas := models.FilterParams{
		SearchTerm: r.FormValue("search"),
		Sizes:      r.Form["size"],
		Types:      r.Form["type"],
		Cities:     r.Form["city"],
		Tags:       r.Form["tag"],
		Status:     r.FormValue("status"),
	}	

	companiesRes,err:=database.ReturnCompaniesByQuery(Paramas)

	if err != nil {
		log.Println("Something went wrong in Searching and Filtering", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sort.Slice(companiesRes, func(i, j int) bool {
		return companiesRes[i].Name < companiesRes[j].Name
	})
	partial := partials.CompaniesList(companiesRes)
	renderErr:=partial.Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ",renderErr)
	}

}
