package handlers

import (
	"net/http"
	"Jordan-Tech-Companies/web/templates/partials"
)

func GetBranchRow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	partials.BranchRow().Render(r.Context(), w)
}

func GetLinkRow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	partials.LinkRow().Render(r.Context(), w)
}
