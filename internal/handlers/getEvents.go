package handlers

import (
	"github.com/m7md2323/JoTech-Directory/internal/database"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
	//"github.com/m7md2323/JoTech-Directory/web/templates/partials"
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
