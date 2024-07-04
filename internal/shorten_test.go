package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LamichhaneBibek/url_shortener/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestHandleShorten_ValidRequest(t *testing.T) {
	// Setup in-memory data structure for testing
	urlDB = make(map[string]models.URL)

	// Prepare valid request body
	requestBody := []byte(`{"url": "https://www.example.com"}`)

	// Create a request
	req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(requestBody))
	assert.NoError(t, err)

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Call the handler function
	HandleShorten(w, req)

	// Assert response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert response content type
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	// Decode the response body
	var response struct {
		ShortURL string `json:"short_url"`
	}
	err = json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)

	// Assert short URL is not empty
	assert.NotEmpty(t, response.ShortURL)

	// Check if short URL exists in urlDB (optional)
	_, ok := urlDB[response.ShortURL]
	assert.True(t, ok)
}

func TestHandleShorten_InvalidRequestBody(t *testing.T) {
	// Prepare invalid request body (missing url field)
	requestBody := []byte(`{}`)

	// Create a request
	req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(requestBody))
	assert.NoError(t, err)

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Call the handler function
	HandleShorten(w, req)

	// Assert response status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert error message in log (optional)
	// assert.Contains(t, log.Messages, "Invalid request body")
}

func TestHandleShorten_InvalidURL(t *testing.T) {
	// Prepare request with invalid URL
	requestBody := []byte(`{"url": "invalid_url"}`)

	// Create a request
	req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(requestBody))
	assert.NoError(t, err)

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Call the handler function
	HandleShorten(w, req)

	// Assert response status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Assert error message in log (optional)
	// assert.Contains(t, log.Messages, "Invalid URL")
}

func TestCreateURL(t *testing.T) {
	originalURL := "https://www.example.com"
	shortURL := createURL(originalURL)

	// Assert short URL is not empty
	assert.NotEmpty(t, shortURL)

	// Check if short URL exists in urlDB (optional)
	// url, ok := urlDB[shortURL]
	// assert.True(t, ok)
	// assert.Equal(t, originalURL, url.LongURL)
}
