package scraper

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type Scraper struct {
	Client *http.Client
	URL    string
}

func NewScraper(url string) *Scraper {
	return &Scraper{
		Client: http.DefaultClient,
		URL:    url,
	}
}

func (s *Scraper) Scrape() (string, string, error) {
	resp, err := s.Client.Get(s.URL)
	if err != nil {
		return "", "", fmt.Errorf("Scraper.Client.Get(%q) returned an error: %s", s.URL, err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("html.Parse(resp.Body) returned an error: %s", err)
	}

	title := findTitle(doc)

	return title, "test description", nil
}
