package handlers

import (
	"github.com/m7md2323/Jordan-Tech-Companies/internal/database"
	"github.com/m7md2323/Jordan-Tech-Companies/web/templates/pages"
	"log"
	"net/http"
)

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	company, err := database.ReturnCompanyByName(name)
	if err != nil {
		log.Println("Error Finding company: ", err)
	}
	err1 := database.DB.Unscoped().Delete(&company).Error
	if err1 != nil {
		log.Println("Error deleting company: ", err1)
	}

	w.WriteHeader(http.StatusOK)
	page := pages.PostAddCompany()
	page.Render(r.Context(), w)

}
