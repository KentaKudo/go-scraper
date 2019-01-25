package scraper

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

type Scraper struct {
	Client *http.Client
	URL    string
}

type OpenGraphAttr = map[string][]string

type Page struct {
	Title         string
	Description   string
	OpenGraphAttr OpenGraphAttr
}

func NewScraper(url string) *Scraper {
	return &Scraper{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
		URL: url,
	}
}

func (s *Scraper) Scrape() (*Page, error) {
	resp, err := s.Client.Get(s.URL)
	if err != nil {
		return nil, fmt.Errorf("Scraper.Client.Get(%q) returned an error: %s", s.URL, err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("html.Parse(resp.Body) returned an error: %s", err)
	}

	p := &Page{
		Title:         findTitle(doc),
		Description:   findDescription(doc),
		OpenGraphAttr: findOpenGraphAttr(doc),
	}

	return p, nil
}

func appendToOpenGraphAttr(oga OpenGraphAttr, property, content string) OpenGraphAttr {
	if v, ok := oga[property]; ok {
		oga[property] = append(v, content)
	} else {
		oga[property] = []string{content}
	}

	return oga
}
