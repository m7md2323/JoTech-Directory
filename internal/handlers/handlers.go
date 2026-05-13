package handlers

import (
	
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/internal/models"
)

func Companies(w http.ResponseWriter,r *http.Request){

	page:=pages.CompaniesList(models.Companies)
	page.Render(r.Context(),w)

}
func CompanyProfile(w http.ResponseWriter,r *http.Request){

	http.NotFound(w,r)

}
func Home(w http.ResponseWriter,r *http.Request){

	http.NotFound(w,r)

}