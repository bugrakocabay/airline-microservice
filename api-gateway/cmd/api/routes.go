package main

import (
	"net/http"

	middleware2 "github.com/bugrakocabay/airline/api-gateway/cmd/middleware"

	"github.com/go-chi/chi/v5"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/auth", app.HandleAuthSubmission)

	mux.Group(func(r chi.Router) {
		r.Use(middleware2.AuthMiddleware)
		r.Mount("/handle", app.handleRouter())
	})

	return mux
}

func (app *Config) handleRouter() http.Handler {
	mux := chi.NewRouter()

	return mux
}
