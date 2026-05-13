package main

import (
	"net/http"
	"log"
	"../../internal/handlers"

)

func main(){

	mux:=http.NewServeMux()

	//Routes

	mux.HandleFunc("/",handlers.Home)
	mux.HandleFunc("/companyProfile",handlers.CompanyProfile)
	mux.HandleFunc("/companies",handlers.Companies)

	log.Fatal(http.ListenAndServe(":8080",nil))

	

}