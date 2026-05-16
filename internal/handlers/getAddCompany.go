package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
)

func GetAddCompany(w http.ResponseWriter, r *http.Request) {

	page := pages.GetAddCompany()
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)

}