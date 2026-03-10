package roles

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/jwt"
)

func Initialize(store database.DBStore, jwtService *jwt.JWTService, log *slog.Logger) http.Handler {
	repo := NewRoleRepository(store)
	svc := NewRoleService(repo, log)
	h := NewRoleHandler(svc, log)
	return NewRoleRouter(jwtService, h)
}