package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Velkommen til Go Webserveren!")
	}

	aboutHandler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Enkel webserver med GoLang")
	}
	
	bjarneHandler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Bjarne")
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/about", aboutHandler )
	http.HandleFunc("/bjarne", bjarneHandler )

	log.Fatal(http.ListenAndServe(":8080", nil))
}