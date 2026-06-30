package handlers

import (
	"net/http"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := pages.Home()
	page.Render(r.Context(), w)
}
func About(w http.ResponseWriter, r *http.Request) {
	page := pages.About()
	page.Render(r.Context(), w)
}
func Contact(w http.ResponseWriter, r *http.Request) {
	page := pages.Contact()
	page.Render(r.Context(), w)
}
