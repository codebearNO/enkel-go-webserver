package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
	
	processHandler := func(w http.ResponseWriter, r *http.Request) {
		go func() {
			start := time.Now()
			var sum int
			for i:=0;i<1000;i++{
				time.Sleep(5*time.Millisecond)
				sum += i
			}
			fmt.Printf("Process took %.4f seconds, and resulted in sum: %v", time.Since(start).Seconds(), sum)
		}()
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/about", aboutHandler )
	http.HandleFunc("/data", dataHandler )
	http.HandleFunc("/process", processHandler)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("C:/Users/Thor/source/repos/go-webserver/static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}