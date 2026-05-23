package handlers

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/web/templates/pages"
	//"Jordan-Tech-Companies/web/templates/partials"
	"fmt"
	"net/http"
)

func Events(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		events,err:=database.ReturnAllEvents()
		if err !=nil {
			fmt.Println("Something went wrong in Events handler",err)
		}
		page := pages.Events(events)
		page.Render(r.Context(), w)
	
}