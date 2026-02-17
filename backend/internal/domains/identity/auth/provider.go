package auth

import (
	"log/slog"
	"net/http"

	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	"github.com/keithyw/pitch-in/internal/domains/identity/users/credentials"
)

func Initialize(store database.DBStore, log *slog.Logger) http.Handler {
	urepo := users.NewUserRepository(store)
	us := users.NewUserService(urepo, log)
	crepo := credentials.NewUserCredentialsRepository(store)
	uc := credentials.NewUserCredentialsService(crepo, log)

	return NewAuthRouter(NewAuthHandler(NewAuthService(us, uc, log), log))
}