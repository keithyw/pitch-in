package permissions

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewPermissionRouter(jwtService *jwt.JWTService, h *PermissionHandler, am *mid.AuthorizationMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.With(am.Authorize(mid.IdentityWrite)).Delete("/{permissionID}", h.Delete)
	r.With(am.Authorize(mid.IdentityRead)).Get("/{permissionID}", h.Get)
	r.With(am.Authorize(mid.IdentityRead)).Get("/", h.FindBy)
	r.With(am.Authorize(mid.IdentityWrite)).Post("/", middleware.DecodeAndValidate(h.Post))
	r.With(am.Authorize(mid.IdentityWrite)).Patch("/{permissionID}", middleware.DecodeAndValidate(h.Patch))
	return r
}