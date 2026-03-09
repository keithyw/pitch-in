package permissions

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/jwt"
)

func Initialize(store database.DBStore, jwtService *jwt.JWTService, log *slog.Logger) http.Handler {
	repo := NewPermissionRepository(store)
	svc := NewPermissionService(repo, log)
	h := NewPermissionHandler(svc, log)
	return NewPermissionRouter(jwtService, h)
}