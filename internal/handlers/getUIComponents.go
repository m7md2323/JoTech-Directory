package handlers

import (
	"net/http"
	"github.com/m7md2323/JoTech-Directory/web/templates/partials"
	"log"
)

func GetBranchRow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	renderErr:=partials.BranchRow().Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ",renderErr)
	}
}

func GetLinkRow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	renderErr:=partials.LinkRow().Render(r.Context(), w)
	if renderErr != nil {
		log.Println("Something went wrong rendering ",renderErr)
	}
}
