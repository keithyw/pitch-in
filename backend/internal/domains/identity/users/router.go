package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewUserRouter(jwtService *jwt.JWTService, h *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.Delete("/{userID}", h.Delete)
	r.Get("/{userID}", h.Get)
	r.Get("/", h.FindBy)
	r.Post("/", middleware.DecodeAndValidate(h.Post))
	r.Patch("/{userID}", middleware.DecodeAndValidate(h.Patch))
	return r
}