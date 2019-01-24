package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KentaKudo/go-scraper/scraper"
)

// TODO
// - Package Layout
// - Open Graph Protocol
// - Timeout
// - Concurrency
// - httptest

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
	type page struct {
		URL                 string              `json:"url"`
		Title               string              `json:"title"`
		Description         string              `json:"description"`
		OpenGraphAttributes map[string][]string `json:"oepn_graph_attributes"`
	}
	ps := []page{}
	for key, values := range queries {
		if key == "url" {
			for _, v := range values {
				p, err := scraper.NewScraper(v).Scrape()
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				ps = append(ps, page{
					URL:                 v,
					Title:               p.Title,
					Description:         p.Description,
					OpenGraphAttributes: p.OpenGraphAttr,
				})
			}
		}
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
