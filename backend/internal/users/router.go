package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewUserRouter(h *UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/", middleware.DecodeAndValidate(h.Post))
	return r
}