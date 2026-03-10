package roles

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewRoleRouter(jwtService *jwt.JWTService, h *RoleHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.Delete("/{roleID}", h.Delete)
	r.Get("/{roleID}", h.Get)
	r.Get("/", h.FindBy)
	r.Post("/", middleware.DecodeAndValidate(h.Post))
	r.Patch("/{roleID}", middleware.DecodeAndValidate(h.Patch))
	return r
}