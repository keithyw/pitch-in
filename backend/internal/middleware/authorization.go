package middleware

import (
	"net/http"

	"github.com/keithyw/pitch-in/pkg/jwt"
	auth "github.com/keithyw/pitch-in/pkg/middleware"
)

type AuthorizationMiddleware struct {
	svc MiddlewareService
}

func NewAuthorizationMiddleware(svc MiddlewareService) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		svc: svc,
	}
}

func (g *AuthorizationMiddleware) Authorize(perm string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(auth.ClaimsKey)
			if val == nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			claims, ok := val.(*jwt.JWTClaim)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			for _, role := range claims.Roles {
				if role == "admin" {
					next.ServeHTTP(w, r)
					return
				}
			}

			isAuth, _ := g.svc.UserHasPermission(claims.UserID, perm)
			if !isAuth {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}