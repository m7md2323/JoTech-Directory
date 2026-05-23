package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/pages"
)

func GetAddEvent(w http.ResponseWriter, r *http.Request) {

	page := pages.AddEvent()
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)

}