package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LamichhaneBibek/url_shortener/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestHandleRedirect_ValidShortCode(t *testing.T) {
	// Setup in-memory data structure for testing
	urlDB = make(map[string]models.URL)

	testURL := "http://localhost:8080/abc123"
	shortCode := "abc123"
	createdAt := time.Now()
	urlDB[shortCode] = models.URL{ID: shortCode, LongURL: testURL, ShortURL: shortCode, CreatedAt: createdAt}

	// Create a request with a valid short code
	req, err := http.NewRequest(http.MethodGet, shortCode, nil)
	assert.NoError(t, err)

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Call the handler function
	HandleRedirect(w, req)

	// Assert response status code (redirect)
	assert.Equal(t, http.StatusFound, w.Code)

	// Assert response location header
	assert.Equal(t, testURL, w.Header().Get("Location"))

	// Assert log message (optional)
	// assert.Contains(t, log.Messages, "Redirecting to original URL")
}

func TestHandleRedirect_InvalidShortCode(t *testing.T) {
	// Setup in-memory data structure for testing
	urlDB = make(map[string]models.URL)

	// Create a request with an invalid short code
	req, err := http.NewRequest(http.MethodGet, "/invalid_code", nil)
	assert.NoError(t, err)

	// Create a recorder to capture the response
	w := httptest.NewRecorder()

	// Call the handler function
	HandleRedirect(w, req)

	// Assert response status code (not found)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Assert error message in log (optional)
	// assert.Contains(t, log.Messages, "URL not found")
}

func TestGetURL_ExistingShortCode(t *testing.T) {
	// Setup in-memory data structure for testing
	urlDB = make(map[string]models.URL)

	testURL := "https://www.example.com"
	shortCode := "xyz789"
	urlDB[shortCode] = models.URL{LongURL: testURL}

	// Get the URL for a valid short code
	url, err := getURL(shortCode)

	// Assert no error returned
	assert.NoError(t, err)

	// Assert retrieved URL matches the expected value
	assert.Equal(t, testURL, url.LongURL)
}

func TestGetURL_NonexistentShortCode(t *testing.T) {
	// Setup in-memory data structure for testing
	urlDB = make(map[string]models.URL)

	// Get the URL for a non-existent short code
	url, err := getURL("nonexistent")

	// Assert expected error returned
	assert.EqualError(t, err, "URL not found")

	// Assert empty URL object returned
	assert.Equal(t, models.URL{}, url)
}
