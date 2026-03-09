package permissions

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewPermissionRouter(jwtService *jwt.JWTService, h *PermissionHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.Delete("/{permissionID}", h.Delete)
	r.Get("/{permissionID}", h.Get)
	r.Get("/", h.FindBy)
	r.Post("/", middleware.DecodeAndValidate(h.Post))
	r.Patch("/{permissionID}", middleware.DecodeAndValidate(h.Patch))
	return r
}