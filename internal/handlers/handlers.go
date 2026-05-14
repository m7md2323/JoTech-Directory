package handlers

import (
	
	"net/http"
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/web/templates/pages"
	//"Jordan-Tech-Companies/internal/models"
)

func Companies(w http.ResponseWriter,r *http.Request){

	page:=pages.CompaniesList(database.ReturnAllCompanies())
	page.Render(r.Context(),w)
	w.WriteHeader(http.StatusOK)

}
func CompanyProfile(w http.ResponseWriter,r *http.Request){

	http.NotFound(w,r)

}
func Home(w http.ResponseWriter,r *http.Request){

	http.NotFound(w,r)

}