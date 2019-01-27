package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealthz(t *testing.T) {
	srv := newServer()
	srv.routes()
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed: %s", "GET", "/healthz", err)
	}

	rr := httptest.NewRecorder()
	srv.router.ServeHTTP(rr, req)

	if got := rr.Code; got != http.StatusOK {
		t.Errorf("HealthCheckHandler returned unexpected status code: want %q, got %q", http.StatusOK, got)
	}

	want := fmt.Sprintln(`{"alive":true}`)
	if rr.Body.String() != want {
		t.Errorf("HealthCheckHandler returned unexpected body: want %q, got %q", want, rr.Body.String())
	}
}
