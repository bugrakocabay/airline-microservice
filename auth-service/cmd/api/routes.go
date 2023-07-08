package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/authenticate", app.authenticate)
	mux.Post("/create", app.createUser)
	return mux
}
