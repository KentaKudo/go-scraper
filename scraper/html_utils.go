package scraper

import "golang.org/x/net/html"

func findTitle(doc *html.Node) string {
	if isTitle(doc) {
		return doc.FirstChild.Data
	}

	for c := doc.FirstChild; c != nil; c = doc.NextSibling {
		t := findTitle(c)
		if t != "" {
			return t
		}
	}

	return ""
}

func isTitle(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}
