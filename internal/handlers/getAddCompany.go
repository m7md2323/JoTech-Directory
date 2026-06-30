package handlers

import (
	"net/http"
	"github.com/m7md2323/Jordan-Tech-Companies/web/templates/pages"
)

func GetAddCompany(w http.ResponseWriter, r *http.Request) {

	page := pages.GetAddCompany()
	page.Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)

}
