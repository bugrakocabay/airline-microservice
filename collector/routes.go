package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (c *Collector) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/collect", c.handleCollection)

	return mux
}
