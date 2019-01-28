package scraper

import (
	"net/http"
	"time"
)

// Factory provides functions to create Scraper
type Factory interface {
	NewScraper(string) *Scraper
}

type factory struct{}

// NewFactory returns an instance which satisfies Facotry interface.
func NewFactory() Factory {
	return &factory{}
}

// NewScraper returns a scraper instance with a given url.
func (f *factory) NewScraper(url string) *Scraper {
	return &Scraper{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
		URL: url,
	}
}
