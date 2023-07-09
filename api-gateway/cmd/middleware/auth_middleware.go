package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/bugrakocabay/airline/api-gateway/cmd/token"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get(authorizationHeaderKey)
		if bearerToken == "" {
			http.Error(w, "authorization header is not provided", http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(bearerToken)
		if len(fields) < 2 {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			http.Error(w, "unsupported authorization type", http.StatusUnauthorized)
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", payload.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
