package api

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/go-chi/chi/v5"
)

// Routes is responsible for routing handlers.
func (app *AuthHandler) Routes() http.Handler {
	r := chi.NewRouter()

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request), method string) {
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		switch method {
		case http.MethodGet:
			r.Get(pattern, handler.ServeHTTP)
		case http.MethodPost:
			r.Post(pattern, handler.ServeHTTP)
		default:
			panic("Unsupported method: " + method)
		}
	}

	// Assuming /create is a POST method and /authenticate is a POST method for this example
	handleFunc("/create", app.createUser, http.MethodPost)
	handleFunc("/authenticate", app.authenticate, http.MethodPost)

	handler := otelhttp.NewHandler(r, "/")

	return handler
}
