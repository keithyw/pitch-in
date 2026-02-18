package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/internal/config"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/identity/auth"
	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewServer(cfg *config.Config, store database.DBStore, log *slog.Logger) http.Handler {
	jwtService := jwt.NewJWTService(cfg.JWTSecretKey, cfg.JWTExpirationTime, log)
	r := chi.NewRouter()
	r.Use(middleware.Cors)
	r.Mount("/auth", auth.Initialize(store, jwtService, log))
	r.Mount("/users", users.Initialize(store, jwtService, log))

	return r

}