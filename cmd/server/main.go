package main

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/handlers"
	"fmt"
	"log"
	"net/http"
	//"Jordan-Tech-Companies/internal/services"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error reading it, using defaults if applicable")
	}
	//initialize the SQLite database
	database.ConnectDatabase()

	//
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//Routes

	//services.AIWebScraper() //only run when needed

	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /about", handlers.About)
	mux.HandleFunc("GET /contact", handlers.Contact)

	mux.HandleFunc("GET /company/{name}", handlers.GetCompanyProfile)
	mux.HandleFunc("GET /companies", handlers.Companies)

	mux.HandleFunc("GET /searchCompanies", handlers.SearchAndFilterCompanies)

	mux.HandleFunc("GET /ui/branch-row", handlers.GetBranchRow)
	mux.HandleFunc("GET /ui/link-row", handlers.GetLinkRow)

	mux.HandleFunc("GET /events", handlers.Events)

	mux.HandleFunc("GET /add_company", handlers.GetAddCompany)
	mux.HandleFunc("POST /add_company", handlers.PostAddCompany)

	mux.HandleFunc("GET /edit_company/{name}", handlers.GetEditCompany)
	mux.HandleFunc("POST /edit_company", handlers.PostEditCompany)

	mux.HandleFunc("POST /delete_company/{name}", handlers.DeleteCompany)

	mux.HandleFunc("GET /add_event", handlers.GetAddEvent)
	mux.HandleFunc("POST /add_event", handlers.PostAddEvent)

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
