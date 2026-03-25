package tags

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
) 

func NewTagRouter(jwtService *jwt.JWTService, h *TagHandler, am *mid.AuthorizationMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.With(am.Authorize(mid.ContentWrite)).Delete("/{tagID}", h.Delete)
	r.With(am.Authorize(mid.ContentRead)).Get("/{tagID}", h.Get)
	r.With(am.Authorize(mid.ContentRead)).Get("/", h.FindBy)
	r.With(am.Authorize(mid.ContentWrite)).Post("/", middleware.DecodeAndValidate(h.Post))
	r.With(am.Authorize(mid.ContentWrite)).Patch("/{tagID}", middleware.DecodeAndValidate(h.Patch))
	r.With(am.Authorize(mid.ContentWrite)).Post("/tagID}/entity", middleware.DecodeAndValidate(h.AttachEntity))
	r.With(am.Authorize(mid.ContentWrite)).Delete("/tagID}/entity", middleware.DecodeAndValidate(h.DetachEntity))
	return r
}