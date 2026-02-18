package users

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/jwt"
)

func Initialize(store database.DBStore, jwtService *jwt.JWTService, log *slog.Logger) http.Handler {
	repo := NewUserRepository(store)
	svc := NewUserService(repo, log)
	h := NewUserHandler(svc, log)
	
	return NewUserRouter(jwtService, h)
}