package api

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

// Routes is responsible for routing handlers.
func (app *AuthHandler) Routes() http.Handler {
	mux := http.NewServeMux()

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	handleFunc("/create", app.createUser)
	handleFunc("/authenticate", app.authenticate)

	handler := otelhttp.NewHandler(mux, "/")

	return handler
}
