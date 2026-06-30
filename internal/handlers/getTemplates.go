package handlers

import (
	"log"
	"net/http"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := pages.Home()
	renderErr := page.Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ", renderErr)
	}
}
func About(w http.ResponseWriter, r *http.Request) {
	page := pages.About()
	renderErr := page.Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ", renderErr)
	}
}
func Contact(w http.ResponseWriter, r *http.Request) {
	page := pages.Contact()
	renderErr := page.Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ", renderErr)
	}
}
