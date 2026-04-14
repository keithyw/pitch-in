package identifier

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewIdentifierRouter(jwtService *jwt.JWTService, h *IdentifierHandler, am *mid.AuthorizationMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.With(am.Authorize(mid.ContentRead)).Post("/", middleware.DecodeMultipart(h.Post))
	return r
}