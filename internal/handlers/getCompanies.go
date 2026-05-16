package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
	"Jordan-Tech-Companies/internal/database"
)

func Companies(w http.ResponseWriter, r *http.Request) {

	page := pages.CompaniesList(database.ReturnAllCompanies())
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)

}