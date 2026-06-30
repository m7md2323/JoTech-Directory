package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/web/templates/pages"
	//"Jordan-Tech-Companies/web/templates/partials"
	"log"
	"net/http"
)

func Events(w http.ResponseWriter, r *http.Request) {
	events, err := database.ReturnAllEvents()
	if err != nil {
		log.Println("Something went wrong in Events handler", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	page := pages.Events(events)
	page.Render(r.Context(), w)
}