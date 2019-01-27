package main

import (
	"log"
	"net/http"
)

// TODO
// - Package Layout
// - Concurrency
// - httptest

func main() {
	srv := newServer()
	srv.routes()
	log.Fatal(http.ListenAndServe(":8080", srv.router))
}
