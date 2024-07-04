package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/LamichhaneBibek/url_shortener/internal/models"
	"github.com/LamichhaneBibek/url_shortener/internal/utils"
)

// In-memory data structure to store URL mappings.
var urlDB = make(map[string]models.URL)

// HandleShorten handles POST requests for "/shorten" and creates a new short URL for the provided long URL.
func HandleShorten(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// If there's an error parsing the request body, return a 400 Bad Request error.
		log.Printf("Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = url.ParseRequestURI(data.URL)
	if err != nil {
		log.Printf("Invalid URL: %s", data.URL)
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	res := createURL(data.URL)

	// Prepare the response with the short URL.
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: res}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Printf("Short URL created: %s", res)
}

// createURL creates a new short URL for the provided long URL, stores it in the urlDB, and returns the short URL.
func createURL(originalURL string) string {
	shortURL := utils.GenerateShortKey(8)

	// Create a new URL object with the short and long URLs and the current timestamp.
	urlDB[shortURL] = models.URL{
		ID:        shortURL,
		LongURL:   originalURL,
		ShortURL:  shortURL,
		CreatedAt: time.Now(),
	}

	// Return the short URL.
	return shortURL
}
