package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age uint8
}

func main(){
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Velkommen til Go Webserveren!")
	}

	aboutHandler := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Enkel webserver med GoLang")
	}
	
	var arne Person = Person{Name: "Arne", Age: 2}

	dataHandler := func(w http.ResponseWriter, r *http.Request){
		
	b, err := json.Marshal(arne)
	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Fprint(w, string(b))
	w.Write(b)
	}
	
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", aboutHandler )
	http.HandleFunc("/data", dataHandler )

	log.Fatal(http.ListenAndServe(":8080", nil))
}