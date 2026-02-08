package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewUserRouter(h *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Delete("/{userID}", h.Delete)
	r.Get("/{userID}", h.Get)
	r.Get("/", h.FindBy)
	r.Post("/", middleware.DecodeAndValidate(h.Post))
	r.Patch("/{userID}", middleware.DecodeAndValidate(h.Patch))
	return r
}