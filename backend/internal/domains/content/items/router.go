package items

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewItemRouter(jwtService *jwt.JWTService, h *ItemHandler, am *mid.AuthorizationMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.With(am.Authorize(mid.ContentWrite)).Delete("/{itemID}", h.Delete)
	r.With(am.Authorize(mid.ContentRead)).Get("/{itemID}", h.Get)
	r.With(am.Authorize(mid.ContentRead)).Get("/", h.FindBy)
	r.With(am.Authorize(mid.ContentWrite)).Post("/", middleware.DecodeAndValidate(h.Post))
	r.With(am.Authorize(mid.ContentWrite)).Patch("/{itemID}", middleware.DecodeAndValidate(h.Patch))
	return r
}