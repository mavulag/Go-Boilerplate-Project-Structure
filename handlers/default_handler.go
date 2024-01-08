package handlers

import "net/http"

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Tri Labs DevOps Limited todo list"))
}
