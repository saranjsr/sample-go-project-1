package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mondoo Engineer!")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second, // Prevents slow client attacks
		WriteTimeout: 10 * time.Second, // Prevents slow server attacks
		IdleTimeout:  60 * time.Second, // Time a connection can remain idle before being closed
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
