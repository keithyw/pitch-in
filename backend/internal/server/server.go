package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keithyw/pitch-in/internal/config"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/assets"
	"github.com/keithyw/pitch-in/internal/domains/content/identifier"
	"github.com/keithyw/pitch-in/internal/domains/content/ingestion"
	"github.com/keithyw/pitch-in/internal/domains/content/items"
	"github.com/keithyw/pitch-in/internal/domains/content/media"
	"github.com/keithyw/pitch-in/internal/domains/identity/auth"
	"github.com/keithyw/pitch-in/internal/domains/identity/permissions"
	"github.com/keithyw/pitch-in/internal/domains/identity/roles"
	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	"github.com/keithyw/pitch-in/internal/domains/taxonomy/tags"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
	"github.com/keithyw/pitch-in/pkg/storage"
	"github.com/keithyw/pitch-in/pkg/templating"
)

func NewServer(cfg *config.Config, store database.DBStore, client storage.StorageClient, log *slog.Logger) http.Handler {
	jwtService := jwt.NewJWTService(cfg.JWTSecretKey, cfg.JWTExpirationTime, log)
	middlewareService := mid.NewMiddlewareService(mid.NewMiddlewareRepository(store), log)
	authorization := mid.NewAuthorizationMiddleware(middlewareService)
	manager := templating.NewTemplateManager(cfg.PromptTemplateDir)

	r := chi.NewRouter()
	r.Use(middleware.Cors)
	r.Mount("/auth", auth.Initialize(store, jwtService, log))
	r.Mount("/assets", assets.Initialize(client, store, jwtService, authorization, log))
	r.Mount("/items", items.Initialize(store, jwtService, authorization, log))
	r.Mount("/permissions", permissions.Initialize(store, jwtService, authorization, log))
	r.Mount("/roles", roles.Initialize(store, jwtService, authorization, log))
	r.Mount("/tags", tags.Initialize(store, jwtService, authorization, log))
	r.Mount("/users", users.Initialize(store, jwtService, authorization, log))
	r.Mount("/identifier", identifier.Initialize(store, jwtService, *manager, authorization, log))
	r.Mount("/ingestion", ingestion.Initialize(client, store, jwtService, authorization, log))
	r.Mount("/media", media.Initialize(store, jwtService, *manager, authorization, log))

	return r

}