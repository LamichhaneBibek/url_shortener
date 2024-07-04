package main

import (
	"log"
	"net/http"

	"github.com/LamichhaneBibek/url_shortener/internal"
)

func main() {
	// Create an HTTP ServeMux to handle requests.
	mux := http.NewServeMux()

	// Register the HandleRedirect function to handle GET requests for "/{shortCode}".
	mux.HandleFunc("GET /{shortCode}", internal.HandleRedirect)

	// Register the HandleShorten function to handle POST requests for "/shorten".
	mux.HandleFunc("POST /shorten", internal.HandleShorten)

	// Log a message indicating that the server has started.
	log.Println("Server started on :8080")

	// Start the HTTP server and listen on port 8080.
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
