package users

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	mid "github.com/keithyw/pitch-in/internal/middleware"
	"github.com/keithyw/pitch-in/pkg/jwt"
)

func Initialize(store database.DBStore, jwtService *jwt.JWTService, authMiddleware *mid.AuthorizationMiddleware, log *slog.Logger) http.Handler {
	repo := NewUserRepository(store)
	svc := NewUserService(repo, log)
	h := NewUserHandler(svc, log)
	
	return NewUserRouter(jwtService, h, authMiddleware)
}