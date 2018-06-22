package main

import (
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

// POST handle로 제한한다.
func PostOnly(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h(w, r)
			return
		}
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
	}
}
