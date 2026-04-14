package identifier

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/templating"
)

func Initialize(store database.DBStore, jwtService *jwt.JWTService, manager templating.TemplateManager, authMiddleware *mid.AuthorizationMiddleware, log *slog.Logger) http.Handler {
	svc := NewIdentifierService(manager, log)
	h := NewIdentifierHandler(svc, log)
	return NewIdentifierRouter(jwtService, h, authMiddleware)
}