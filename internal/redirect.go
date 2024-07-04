package internal

import (
	"errors"
	"log"
	"net/http"

	"github.com/LamichhaneBibek/url_shortener/internal/models"
)

// HandleRedirect handles GET requests for "/{shortCode}" and redirects the user to the original URL.
func HandleRedirect(w http.ResponseWriter, r *http.Request) {

	// Extract the shortCode from the request path.
	id := r.URL.Path
	url, err := getURL(id)
	if err != nil {

		// If the URL is not found, return a 404 Not Found error.
		log.Printf("URL not found: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	log.Printf("Redirecting to original URL: %s", url.LongURL)
	// Redirect the user to the original URL with a 302 Found status code.
	http.Redirect(w, r, url.LongURL, http.StatusFound)
}

// getURL retrieves the URL object from the urlDB using the given ID.
// Returns an error if the URL is not found.
func getURL(id string) (models.URL, error) {

	url, ok := urlDB[id]
	if !ok {
		log.Printf("URL not found")
		return models.URL{}, errors.New("URL not found")
	}
	return url, nil
}
