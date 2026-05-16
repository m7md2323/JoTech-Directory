package main

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/handlers"
	"fmt"
	"log"
	"net/http"
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

	mux.HandleFunc("GET /{$}", handlers.Home)
	//mux.HandleFunc("GET /company/{id}", handlers.Company)
	mux.HandleFunc("GET /companies", handlers.Companies)
	mux.HandleFunc("GET /about", handlers.About)
	mux.HandleFunc("GET /add_company", handlers.GetAddCompany)
	mux.HandleFunc("POST /add_company", handlers.PostAddCompany)
	mux.HandleFunc("GET /contact", handlers.Contact)

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
