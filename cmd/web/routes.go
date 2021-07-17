package main

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/sztyui/ws/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))

	return mux
}
