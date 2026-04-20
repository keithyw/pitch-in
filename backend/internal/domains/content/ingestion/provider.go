package ingestion

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/assets"
	"github.com/keithyw/pitch-in/internal/domains/content/items"
	"github.com/keithyw/pitch-in/internal/domains/taxonomy/tags"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/storage"
)

func Initialize(client storage.StorageClient, store database.DBStore, jwtService *jwt.JWTService, authMiddleware *mid.AuthorizationMiddleware, log *slog.Logger) http.Handler {
	assetRepo := assets.NewAssetRepository(store)
	itemRepo := items.NewItemRepository(store)
	tagRepo := tags.NewTagRepository(store)
	assetService := assets.NewAssetService(assetRepo, client, log)
	itemService := items.NewItemService(itemRepo, tagRepo, log)
	tagService := tags.NewTagService(tagRepo, log)
	h := NewIngestionHandler(assetService, itemService, tagService, log)
	return NewIngestionRouter(jwtService, h, authMiddleware)
}