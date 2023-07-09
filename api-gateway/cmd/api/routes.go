package main

import (
	middleware2 "github.com/bugrakocabay/airline/api-gateway/cmd/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Group(func(r chi.Router) {
		r.Use(middleware2.AuthMiddleware)
		r.Mount("/handle", app.handleRouter())
	})

	mux.Post("/auth", app.HandleAuthSubmission)

	return mux
}

func (app *Config) handleRouter() http.Handler {
	mux := chi.NewRouter()

	return mux
}
