package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Routes is responsible for routing handlers.
func (app *AuthHandler) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/create", app.createUser)
	mux.Post("/authenticate", app.authenticate)

	return mux
}
