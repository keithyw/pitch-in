package assets

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/storage"
)

func Initialize(client storage.StorageClient, store database.DBStore, jwtService *jwt.JWTService, authMiddleware *mid.AuthorizationMiddleware, log *slog.Logger) http.Handler {
	repo := NewAssetRepository(store)
	svc := NewAssetService(repo, client, log)
	h := NewAssetHandler(svc, log)
	return NewAssetRouter(jwtService, h, authMiddleware)
}