package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {

	page := pages.Home()
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)

}
func About(w http.ResponseWriter, r *http.Request) {

	page := pages.About()
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)

}
func Contact(w http.ResponseWriter, r *http.Request) {

	page := pages.Contact()
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)

}