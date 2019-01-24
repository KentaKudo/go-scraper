package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScrape(t *testing.T) {
	cases := [...]struct {
		title string
		desc  string
		input string
	}{
		{
			title: "",
			desc:  "",
			input: "",
		},
		{
			title: "test title",
			desc:  "",
			input: `
<html>
	<head><title>test title</title></head>
	<body></body>
</html>`,
		},
		{
			title: "",
			desc:  "test description",
			input: `
<html>
	<head><meta name="description" content="test description"></head>
	<body></body>
</html>`,
		},
		{
			title: "test title",
			desc:  "test description",
			input: `
<html>
	<head>
		<title>test title</title>
		<meta name="description" content="test description">
	</head>
	<body></body>
</html>`,
		},
	}
	for _, c := range cases {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(c.input))
		}))

		sut := &Scraper{
			Client: srv.Client(),
			URL:    srv.URL,
		}

		p, err := sut.Scrape()
		if err != nil {
			t.Errorf("Scrape() returned unexpected error: %s", err)
		}
		if p.Title != c.title {
			t.Errorf("Unexpected title returned: want %q, got %q", c.title, p.Title)
		}
		if p.Description != c.desc {
			t.Errorf("Unexpected description returned: want %q, got %q", c.desc, p.Description)
		}

		srv.Close()
	}
}
