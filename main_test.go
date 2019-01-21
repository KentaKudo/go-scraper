package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	method := "GET"
	endpoint := "/healthz"
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed: %s", method, endpoint, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	cases := []struct {
		status int
		want   string
	}{
		{status: http.StatusOK, want: fmt.Sprintln(`{"alive":true}`)},
	}
	for _, c := range cases {
		handler.ServeHTTP(rr, req)
		if got := rr.Code; got != c.status {
			t.Errorf("HealthCheckHandler returned unexpected status code: want %q, got %q", c.status, got)
		}

		if rr.Body.String() != c.want {
			t.Errorf("HealthCheckHandler returned unexpected body: want %q, got %q", c.want, rr.Body.String())
		}
	}
}

func TestScraperHandler(t *testing.T) {
	method := "GET"
	endpoint := "/"
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed: %s", method, endpoint, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ScraperHandler)

	handler.ServeHTTP(rr, req)
	if got := rr.Code; got != http.StatusOK {
		t.Errorf("ScraperHandler returned unexpected status code: want %q, got %q", http.StatusOK, got)
	}
}

func TestScraperHandlerParseQuery(t *testing.T) {
	method := "GET"
	endpoint := "/"
	query := url.Values{"foo": []string{"bar"}}.Encode()
	url := fmt.Sprintf("%s?%s", endpoint, query)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed: %s", method, url, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ScraperHandler)

	handler.ServeHTTP(rr, req)
	if got := rr.Code; got != http.StatusOK {
		t.Errorf("ScraperHandler returned unexpected status code: want %q, got %q", http.StatusOK, got)
	}

	want := fmt.Sprintln("key: foo, value: bar")
	if rr.Body.String() != want {
		t.Errorf("ScraperHandler returned unexpected body: want %q, got %q", want, rr.Body.String())
	}
}
