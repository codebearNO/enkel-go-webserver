package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  uint8
}

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Velkommen til Go Webserveren!")
	}

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Enkel webserver med GoLang")
	}

	var arne Person = Person{Name: "Arne", Age: 2}

	dataHandler := func(w http.ResponseWriter, r *http.Request) {
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
			for i := 0; i < 1000; i++ {
				time.Sleep(5 * time.Millisecond)
				sum += i
			}
			fmt.Printf("Process took %.4f seconds, and resulted in sum: %v", time.Since(start).Seconds(), sum,)
			fmt.Print("\n")
		}()
	}

	mux.HandleFunc("/", handler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/process", processHandler)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("C:/Users/Thor/source/repos/go-webserver/static"))))

	// Start the HTTP server in a separate goroutine
	go func() {
		fmt.Println("Starting HTTP server...")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server error: %s\n", err)
		}
	}()

	// Wait for the context to be canceled
	select {
	case <-ctx.Done():
		// Shutdown the server gracefully
		fmt.Println("Shutting down HTTP server gracefully...")
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelShutdown()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			fmt.Printf("HTTP server shutdown error: %s\n", err)
		}
	}

	fmt.Println("HTTP server stopped.")
}
