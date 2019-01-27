package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type server struct {
	router *http.ServeMux
}

func newServer() *server {
	return &server{router: http.DefaultServeMux}
}

func ok(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		ng(w, r, err, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if _, err := buf.WriteTo(w); err != nil {
		fmt.Fprintf(os.Stderr, "error in writing from buffer to response: %q", buf.String())
	}
}

func ng(w http.ResponseWriter, r *http.Request, err error, code int) {
	errObj := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(errObj); err != nil {
		fmt.Fprintf(os.Stderr, "error in writing from buffer to response: %q", errObj.Error)
	}
}
