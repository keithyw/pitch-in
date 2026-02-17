package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/identity/auth"
	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewServer(store database.DBStore, log *slog.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Cors)
	r.Mount("/auth", auth.Initialize(store, log))
	r.Mount("/users", users.Initialize(store, log))

	return r

}