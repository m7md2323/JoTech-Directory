package main

import (
	"Jordan-Tech-Companies/internal/database"
	"Jordan-Tech-Companies/internal/handlers"
	"fmt"
	"log"
	"net/http"
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
	mux.HandleFunc("GET /about", handlers.About)

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
