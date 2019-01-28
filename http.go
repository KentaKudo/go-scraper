package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/KentaKudo/go-scraper/scraper"
)

type server struct {
	router         *http.ServeMux
	ScraperFactory scraper.Factory
}

func newServer() *server {
	return &server{
		router:         http.DefaultServeMux,
		ScraperFactory: scraper.NewFactory(),
	}
}

// OK outputs 2XX result in json.
func OK(w http.ResponseWriter, v interface{}, code int) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		Error(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if _, err := buf.WriteTo(w); err != nil {
		fmt.Fprintf(os.Stderr, "error in writing from buffer to response: %q", buf.String())
	}
}

// Error outputs error HTTP response in json. Similar to http.Error()
func Error(w http.ResponseWriter, err error, code int) {
	errObj := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(errObj); err != nil {
		fmt.Fprintf(os.Stderr, "error in encoding error object: %q", errObj.Error)
	}
}
