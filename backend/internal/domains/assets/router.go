package assets

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func NewAssetRouter(jwtService *jwt.JWTService, h *AssetHandler, am *mid.AuthorizationMiddleware) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(jwtService))
	r.With(am.Authorize(mid.AssetWrite)).Delete("/{assetID}", h.Delete)
	r.With(am.Authorize(mid.AssetRead)).Get("/{assetID}", h.Get)
	r.With(am.Authorize(mid.AssetRead)).Get("/", h.FindBy)
	r.With(am.Authorize(mid.AssetWrite)).Post("/", middleware.DecodeMultipart(h.Post))
	r.With(am.Authorize(mid.AssetWrite)).Patch("/{assetID}", middleware.DecodeMultipart(h.Patch))
	return r
}