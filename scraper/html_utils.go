package scraper

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func findTitle(doc *html.Node) string {
	if isTitle(doc) {
		return doc.FirstChild.Data
	}

	return traverse(doc, findTitle)
}

func isTitle(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func findDescription(doc *html.Node) string {
	if isDescription(doc) {
		return attributes(doc.Attr).Content()
	}

	return traverse(doc, findDescription)
}

func isDescription(n *html.Node) bool {
	return n.Type == html.ElementNode &&
		n.DataAtom == atom.Meta &&
		n.Data == "meta" &&
		attributes(n.Attr).isDescription()
}

func findOpenGraphAttr(doc *html.Node) OpenGraphAttr {
	return findOpenGraphAttrWithAggr(doc, OpenGraphAttr{})
}

func findOpenGraphAttrWithAggr(n *html.Node, aggr OpenGraphAttr) OpenGraphAttr {
	if isOpenGraphAttr(n) {
		as := attributes(n.Attr)
		return appendToOpenGraphAttr(aggr, as.Property(), as.Content())
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		aggr = findOpenGraphAttrWithAggr(c, aggr)
	}

	return aggr
}

func isOpenGraphAttr(n *html.Node) bool {
	return n.Type == html.ElementNode &&
		n.DataAtom == atom.Meta &&
		n.Data == "meta" &&
		attributes(n.Attr).isOpenGraphAttr()
}

func traverse(n *html.Node, f func(*html.Node) string) string {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str := f(c)
		if str != "" {
			return str
		}
	}
	return ""
}

type attributes []html.Attribute

func (as attributes) isDescription() bool {
	for _, a := range as {
		if a.Key == "name" && a.Val == "description" {
			return true
		}
	}
	return false
}

func (as attributes) isOpenGraphAttr() bool {
	for _, a := range as {
		if a.Key == "property" && strings.HasPrefix(a.Val, "og:") {
			return true
		}
	}
	return false
}

func (as attributes) Content() string {
	for _, a := range as {
		if a.Key == "content" {
			return a.Val
		}
	}

	return ""
}

func (as attributes) Property() string {
	for _, a := range as {
		if a.Key == "property" {
			return a.Val
		}
	}
	return ""
}
