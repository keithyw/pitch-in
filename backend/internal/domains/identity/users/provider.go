package users

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
)

func Initialize(store database.DBStore, log *slog.Logger) http.Handler {
	repo := NewUserRepository(store)
	svc := NewUserService(repo, log)
	h := NewUserHandler(svc, log)
	
	return NewUserRouter(h)
}