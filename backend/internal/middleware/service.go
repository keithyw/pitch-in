package middleware

import (
	"fmt"
	"log/slog"
)

type MiddlewareService interface {
	UserHasPermission(userID int64, code string) (bool, error)
}

type middlewareServiceImpl struct {
	repository MiddlewareRepository
	log *slog.Logger
}

func NewMiddlewareService(repo MiddlewareRepository, log *slog.Logger) MiddlewareService {
	return &middlewareServiceImpl{
		repository: repo,
		log: log,
	}
}

func (s *middlewareServiceImpl) UserHasPermission(userID int64, code string) (bool, error) {
	hasPerm, err := s.repository.UserHasPermission(userID, code)
	if err != nil {
		s.log.Error("Faile checking if user has permission", "userID", userID, "error", err)
		return false, fmt.Errorf("User has permission error: %w", err)
	}
	return hasPerm, nil
}