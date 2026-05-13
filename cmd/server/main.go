package main

import (
	"log"
	"net/http"

	//"github.com/m7md2323/Jordan-Tech-Companies/internal/handlers"
	"Jordan-Tech-Companies/internal/handlers"
)

func main() {

	mux := http.NewServeMux()

	//Routes

	mux.HandleFunc("GET /", handlers.Home)
	mux.HandleFunc("GET /companyProfile", handlers.CompanyProfile)
	mux.HandleFunc("GET /companies", handlers.Companies)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
