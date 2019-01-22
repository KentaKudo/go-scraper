package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KentaKudo/go-scraper/scraper"
)

func main() {
	http.HandleFunc("/healthz", HealthCheckHandler)
	http.HandleFunc("/", ScraperHandler)
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
	queries := r.URL.Query()
	type result struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	var re result
	for key, values := range queries {
		if key == "url" {
			for _, v := range values {
				t, d, err := scraper.NewScraper(v).Scrape()
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				re = result{Title: t, Description: d}
				break
			}
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(re)
}
