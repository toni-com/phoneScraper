package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// simple mock HTML page
const mockHTML = `
<!DOCTYPE html>
<html>
<body>
	<h1>Product Page</h1>
	<div class="product-price">1,200.50</div>
	<span id="out-of-stock">Sold Out</span>
</body>
</html>
`

func TestScrape(t *testing.T) {
	// set up a local test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(mockHTML))
	}))
	defer server.Close()

	want := "1,200.50"
	selector := ".product-price"

	got, err := Scrape(server.URL, selector)

	// 4. Assertions
	if err != nil {
		t.Fatalf("Scrape failed unexpectedly: %v", err)
	}

	if got != want {
		t.Errorf("Scraper extraction mismatch. Want '%s', got '%s'", want, got)
	}
}

func TestScrapeNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(mockHTML))
	}))
	defer server.Close()

	// This class does not exist in our mock HTML
	selector := ".non-existent-class"

	got, err := Scrape(server.URL, selector)

	if err != nil {
		t.Fatalf("Scrape error: %v", err)
	}

	// If selector isn't found, we expect an empty string (or handle it how you prefer)
	if got != "" {
		t.Errorf("Expected empty string for missing selector, got '%s'", got)
	}
}
