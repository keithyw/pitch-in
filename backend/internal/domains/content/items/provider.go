package items

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/taxonomy/tags"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
)

func Initialize(store database.DBStore, jwtService *jwt.JWTService, authMiddleware *mid.AuthorizationMiddleware, log *slog.Logger) http.Handler {
	repo := NewItemRepository(store)
	tagRepo := tags.NewTagRepository(store)
	svc := NewItemService(repo, tagRepo, log)
	h := NewItemHandler(svc, log)
	return NewItemRouter(jwtService, h, authMiddleware)
}