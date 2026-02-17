package auth

import (
	"fmt"
	"log/slog"

	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	"github.com/keithyw/pitch-in/internal/domains/identity/users/credentials"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (*users.User, error)
	Logout() (error)
	PasswordReset() (error)
	Register(u users.User, password string) (*users.User, error)
}

type authServiceImpl struct {
	userService users.UserService
	credentialsService credentials.UserCredentialsService
	log *slog.Logger
}

func NewAuthService(us users.UserService, uc credentials.UserCredentialsService, log *slog.Logger) AuthService {
	return &authServiceImpl{
		userService: us,
		credentialsService: uc,
		log: log,
	}
}

func (s *authServiceImpl) Login(email, password string) (*users.User, error) {
	u, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	creds, err := s.credentialsService.GetUserCredentials(u.ID)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(password)); err != nil {
		s.log.Error("Failed password comparison", "error", err)
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}
	return u, nil
}

func (s *authServiceImpl) Logout() (error) {
	return nil
}

func (s *authServiceImpl) PasswordReset() (error) {
	return nil
}

func (s *authServiceImpl) Register(u users.User, password string) (*users.User, error) {
	newUser, err := s.userService.CreateUser(u)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("Failed generating hash", "error", err)
		return nil, fmt.Errorf("Failed generating passwordhash: %w", err)
	}

	creds := credentials.UserCredentials{
		UserID: newUser.ID,
		PasswordHash: string(hash),
	}

	_, err = s.credentialsService.CreateUserCredentials(creds)
	if err != nil {
		return nil, err
	}
	
	return newUser, nil
}