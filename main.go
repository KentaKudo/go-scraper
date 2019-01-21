package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/KentaKudo/go-scraper/scraper"
)

func main() {
	http.HandleFunc("/healthz", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	h := struct {
		Alive bool `json:"alive"`
	}{Alive: true}

	if err := json.NewEncoder(w).Encode(h); err != nil {
		// log error
		http.Error(w, err.Error(), 500)
		return
	}
}

func ScraperHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%q", r.RequestURI)
	queries := r.URL.Query()
	for key, values := range queries {
		for _, v := range values {
			fmt.Fprintf(w, "key: %s, value: %s\n", key, v)
		}
	}
}
