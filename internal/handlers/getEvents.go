package handlers

import (
	"github.com/m7md2323/Jordan-Tech-Companies/internal/database"
	"github.com/m7md2323/Jordan-Tech-Companies/web/templates/pages"
	//"github.com/m7md2323/Jordan-Tech-Companies/web/templates/partials"
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
