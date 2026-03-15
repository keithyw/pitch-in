package roles

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewRoleRouter(jwtService *jwt.JWTService, h *RoleHandler, am *mid.AuthorizationMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.With(am.Authorize(mid.IdentityWrite)).Delete("/{roleID}", h.Delete)
	r.With(am.Authorize(mid.IdentityRead)).Get("/{roleID}", h.Get)
	r.With(am.Authorize(mid.IdentityRead)).Get("/", h.FindBy)
	r.With(am.Authorize(mid.IdentityWrite)).Post("/", middleware.DecodeAndValidate(h.Post))
	r.With(am.Authorize(mid.IdentityWrite)).Patch("/{roleID}", middleware.DecodeAndValidate(h.Patch))
	r.With(am.Authorize(mid.IdentityWrite)).Post("/{roleID}/permissions", middleware.DecodeAndValidate(h.AttachPermission))
	r.With(am.Authorize(mid.IdentityWrite)).Delete("/{roleID}/permissions/{permissionID}", h.DetachPermission)
	return r
}