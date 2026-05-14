package main

import (
	"fmt"
	"log"
	"net/http"
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/handlers"
)

func main() {

	//initialize the SQLite database
	database.ConnectDatabase()

	
	//
	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("./web/static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//Routes

	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /companyProfile", handlers.CompanyProfile)
	mux.HandleFunc("GET /companies", handlers.Companies)

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

	
}
