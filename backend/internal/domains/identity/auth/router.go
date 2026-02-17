package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewAuthRouter(h *AuthHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/login", middleware.DecodeAndValidate(h.Login))
	r.Post("/register", middleware.DecodeAndValidate(h.Register))
	return r
}