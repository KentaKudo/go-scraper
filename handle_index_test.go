package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScraperHandler(t *testing.T) {
	sut := newServer()
	sut.routes()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed: %s", "GET", "/", err)
	}

	rr := httptest.NewRecorder()
	sut.router.ServeHTTP(rr, req)

	if got := rr.Code; got != http.StatusOK {
		t.Errorf("ScraperHandler returned unexpected status code: want %q, got %q", http.StatusOK, got)
	}
}
