package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/keithyw/pitch-in/pkg/jwt"
)

type contextKey string
const ClaimsKey contextKey = "claims"

func AuthMiddleware(s *jwt.JWTService) func(http.Handler)  http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == ""  {
				http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(header, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization Format", http.StatusUnauthorized)
				return
			}

			claims, err := s.ParseJWT(parts[1])
			if err != nil {
				http.Error(w, "Invalid or Expired Token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}