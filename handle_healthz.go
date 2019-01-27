package main

import (
	"net/http"
)

func (s *server) handleHealthz() http.HandlerFunc {
	type response struct {
		Alive bool `json:"alive"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		res := response{Alive: true}
		ok(w, r, res, http.StatusOK)
	}
}
