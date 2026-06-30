package handlers

import (
	"net/http"
	"github.com/m7md2323/JoTech-Directory/web/templates/pages"
	"log"
)

func GetAddCompany(w http.ResponseWriter, r *http.Request) {
	
	w.WriteHeader(http.StatusOK)
	page := pages.GetAddCompany()
	renderErr:=page.Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ",renderErr)
	}


}
