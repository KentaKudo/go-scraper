package main

import (
	"net/http"

	"github.com/KentaKudo/go-scraper/scraper"
)

func (s *server) handleIndex() http.HandlerFunc {
	type response struct {
		URL                 string              `json:"url"`
		Title               string              `json:"title"`
		Description         string              `json:"description"`
		OpenGraphAttributes map[string][]string `json:"oepn_graph_attributes"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		rs := []response{}
		for key, values := range qs {
			if key == "url" {
				for _, v := range values {
					p, err := scraper.NewScraper(v).Scrape()
					if err != nil {
						ng(w, r, err, http.StatusBadRequest)
						return
					}
					rs = append(rs, response{
						URL:                 v,
						Title:               p.Title,
						Description:         p.Description,
						OpenGraphAttributes: p.OpenGraphAttr,
					})
				}
			}
		}

		ok(w, r, rs, http.StatusOK)
	}
}
