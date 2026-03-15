package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewUserRouter(jwtService *jwt.JWTService, h *UserHandler, am *mid.AuthorizationMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.With(am.Authorize(mid.IdentityWrite)).Delete("/{userID}", h.Delete)
	r.With(am.Authorize(mid.IdentityRead)).Get("/{userID}", h.Get)
	r.With(am.Authorize(mid.IdentityRead)).Get("/", h.FindBy)
	r.With(am.Authorize(mid.IdentityWrite)).Post("/", middleware.DecodeAndValidate(h.Post))
	r.With(am.Authorize(mid.IdentityWrite)).Patch("/{userID}", middleware.DecodeAndValidate(h.Patch))
	r.With(am.Authorize(mid.IdentityWrite)).Post("/{userID}/roles", middleware.DecodeAndValidate(h.AttachRole))
	r.With(am.Authorize(mid.IdentityWrite)).Delete("/{userID}/roles/{roleID}", h.DetachRole)
	return r
}