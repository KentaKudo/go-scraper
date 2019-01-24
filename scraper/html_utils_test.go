package scraper

import (
	"bytes"
	"reflect"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestFindTitle(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{input: "", want: ""},
		{input: `<html><head></head><body></body></html>`, want: ""},
		{input: `<html><head><title>test</title></head><body></body></html>`, want: "test"},
		{input: `<html><head><meta name="description"></head><body></body></html>`, want: ""},
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
				DataAtom: atom.Title,
				Data:     "title",
			},
			want: true,
		},
		{
			input: &html.Node{
				Type:     html.DoctypeNode,
				DataAtom: atom.Title,
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

func TestFindDescription(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{input: "", want: ""},
		{input: `<html><head></head><body></body></html>`, want: ""},
		{input: `<html><head><title>test</title></head><body></body></html>`, want: ""},
		{input: `<html><head><meta name="description" content="test"></head><body></body></html>`, want: "test"},
	}
	for _, c := range cases {
		r := bytes.NewReader([]byte(c.input))
		doc, err := html.Parse(r)
		if err != nil {
			t.Errorf("html.Parse(%q) returned an unexpected error: %s", c.input, err)
		}

		if got := findDescription(doc); got != c.want {
			t.Errorf("findDescription(%q): got %q, want %q", c.input, got, c.want)
		}
	}
}

func TestIsDescription(t *testing.T) {
	cases := []struct {
		input *html.Node
		want  bool
	}{
		{
			input: &html.Node{
				Type:     html.ElementNode,
				DataAtom: atom.Meta,
				Data:     "meta",
				Attr: []html.Attribute{
					{Namespace: "", Key: "name", Val: "description"},
					{Namespace: "", Key: "content", Val: "test"},
				},
			},
			want: true,
		},
		{
			input: &html.Node{
				Type:     html.DoctypeNode,
				DataAtom: atom.Meta,
				Data:     "meta",
				Attr: []html.Attribute{
					{Namespace: "", Key: "name", Val: "description"},
					{Namespace: "", Key: "content", Val: "test"},
				},
			},
			want: false,
		},
		{
			input: &html.Node{
				Type:     html.DoctypeNode,
				DataAtom: atom.Div,
				Data:     "meta",
				Attr: []html.Attribute{
					{Namespace: "", Key: "name", Val: "description"},
					{Namespace: "", Key: "content", Val: "test"},
				},
			},
			want: false,
		},
		{
			input: &html.Node{
				Type:     html.DoctypeNode,
				DataAtom: atom.Meta,
				Data:     "div",
				Attr: []html.Attribute{
					{Namespace: "", Key: "name", Val: "description"},
					{Namespace: "", Key: "content", Val: "test"},
				},
			},
			want: false,
		},
		{
			input: &html.Node{
				Type:     html.DoctypeNode,
				DataAtom: atom.Meta,
				Data:     "meta",
				Attr: []html.Attribute{
					{Namespace: "", Key: "name", Val: "author"},
					{Namespace: "", Key: "content", Val: "John Doe"},
				},
			},
			want: false,
		},
	}
	for _, c := range cases {
		if got := isDescription(c.input); got != c.want {
			t.Errorf("isDescription(%v): got %t, want %t", c.input, got, c.want)
		}
	}
}

func TestFindOpenGraphAttr(t *testing.T) {
	cases := []struct {
		input string
		want  map[string][]string
	}{
		{input: "", want: map[string][]string{}},
		{input: `<html><head></head><body></body></html>`, want: map[string][]string{}},
		{input: `<html><head><title>test</title></head><body></body></html>`, want: map[string][]string{}},
		{input: `<html><head><meta property="og:type" content="test"></head><body></body></html>`, want: map[string][]string{"og:type": []string{"test"}}},
		{input: `<html><head><meta property="unknown" content="test"></head><body></body></html>`, want: map[string][]string{}},
		{
			input: `<html><head>
				<meta property="og:type1" content="test"><meta property="og:type2" content="test">
				</head><body></body></html>`,
			want: map[string][]string{"og:type1": []string{"test"}, "og:type2": []string{"test"}},
		},
		{
			input: `<html><head>
				<meta property="og:type" content="test1"><meta property="og:type" content="test2">
				</head><body></body></html>`,
			want: map[string][]string{"og:type": []string{"test1", "test2"}},
		},
	}
	for _, c := range cases {
		r := bytes.NewReader([]byte(c.input))
		doc, err := html.Parse(r)
		if err != nil {
			t.Errorf("html.Parse(%q) returned an unexpected error: %s", c.input, err)
		}

		if got := findOpenGraphAttr(doc); !reflect.DeepEqual(got, c.want) {
			t.Errorf("findOpenGraphAttr(%q): got %q, want %q", c.input, got, c.want)
		}
	}
}
