package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/internal/config"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/identity/auth"
	"github.com/keithyw/pitch-in/internal/domains/identity/permissions"
	"github.com/keithyw/pitch-in/internal/domains/identity/roles"
	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewServer(cfg *config.Config, store database.DBStore, log *slog.Logger) http.Handler {
	jwtService := jwt.NewJWTService(cfg.JWTSecretKey, cfg.JWTExpirationTime, log)
	middlewareService := mid.NewMiddlewareService(mid.NewMiddlewareRepository(store), log)
	authorization := mid.NewAuthorizationMiddleware(middlewareService)
	r := chi.NewRouter()
	r.Use(middleware.Cors)
	r.Mount("/auth", auth.Initialize(store, jwtService, log))
	r.Mount("/permissions", permissions.Initialize(store, jwtService, authorization, log))
	r.Mount("/roles", roles.Initialize(store, jwtService, authorization, log))
	r.Mount("/users", users.Initialize(store, jwtService, authorization, log))

	return r

}