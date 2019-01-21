package scraper

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestFindTitle(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{input: `<html><head></head><body></body></html>`, want: ""},
		{input: `<html><head><title>test</title></head><body></body></html>`, want: "test"},
	}
	for _, c := range cases {
		r := bytes.NewReader([]byte(c.input))
		doc, err := html.Parse(r)
		if err != nil {
			t.Errorf("html.Parse(%q) returned an unexpected error: %s", c.input, err)
		}

		if got := findTitle(doc); got != c.want {
			t.Errorf("findTitle(%q): got %q, want %q", c.input, got, c.want)
		}
	}
}

func TestIsTitle(t *testing.T) {
	cases := []struct {
		input *html.Node
		want  bool
	}{
		{
			input: &html.Node{
				Type:     html.ElementNode,
				DataAtom: atom.Lookup([]byte("title")),
				Data:     "title",
			},
			want: true,
		},
		{
			input: &html.Node{
				Type:     html.DoctypeNode,
				DataAtom: atom.Lookup([]byte("title")),
				Data:     "title",
			},
			want: false,
		},
		{
			input: &html.Node{
				Type:     html.ElementNode,
				DataAtom: atom.Lookup([]byte("test")),
				Data:     "div",
			},
			want: false,
		},
	}
	for _, c := range cases {
		if got := isTitle(c.input); got != c.want {
			t.Errorf("isTitle(%v): got %t, want %t", c.input, got, c.want)
		}
	}
}
