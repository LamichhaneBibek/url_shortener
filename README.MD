URL Shortener API Documentation
==============================

This document provides an overview of the URL Shortener API implemented in Go.

API Endpoints
-------------

### 1. Shorten URL

- **Endpoint:** `POST /shorten`
- **Request Body:** JSON object with a single property `url` containing the long URL to be shortened.
- **Response:** JSON object with a single property `short_url` containing the shortened URL.

Example request:
```json
{
  "url": "https://www.example.com/long-url"
}
```
Example response:
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```
### 2. Redirect to Original URL

- **Endpoint:** `GET /{shortCode}`
- **Path Parameter:** `shortCode` - The shortened URL code.
- **Response:** Redirects the user to the original URL with a 302 Found status code.

Example: `GET http://localhost:8080/abc123` will redirect to `https://www.example.com/long-url`.

Implementation Details
---------------------

The API is implemented using the Go programming language and the net/http package. The following packages and files are included in the project:

- `cmd/main.go`: Contains the main function that sets up the HTTP server and registers the request handlers.
- `internal/models/url.go`: Defines the URL data structure used to store shortened URLs.
- `internal/utils/utils.go`: Provides utility functions for generating random short keys.
- `internal/redirect.go`: Handles GET requests for redirecting short URLs to their original URLs.
- `internal/shorten.go`: Handles POST requests for shortening URLs.

The URLs are stored in an in-memory data structure (map) for simplicity. 

Requirements
------------

- Go 1.16 or higher

Build and Run
-------------

1. Clone the repository: `git clone https://github.com/LamichhaneBibek/url_shortener.git`
2. Navigate to the project directory: `cd url_shortener`
3. Build and run the application: `go run cmd/main.go`
4. The API will be available at `http://localhost:8080`.

Testing
-------

You can test the API using tools like curl, Postman, or any HTTP client of your choice.