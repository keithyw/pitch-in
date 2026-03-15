package roles

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
)

func Initialize(store database.DBStore, jwtService *jwt.JWTService, authMiddleware *mid.AuthorizationMiddleware, log *slog.Logger) http.Handler {
	repo := NewRoleRepository(store)
	svc := NewRoleService(repo, log)
	h := NewRoleHandler(svc, log)
	return NewRoleRouter(jwtService, h, authMiddleware)
}