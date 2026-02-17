package credentials

import (
	"fmt"
	"log/slog"
)

type UserCredentialsService interface {
	CreateUserCredentials(uc UserCredentials) (*UserCredentials, error)
	DeleteUserCredentials(userId int64) error
	GetUserCredentials(userId int64) (*UserCredentials, error)
	UpdateUserCredentials(uc UserCredentials) (*UserCredentials, error)
}

type userCredentialsServiceImpl struct {
	repository UserCredentialsRepository
	log *slog.Logger
}

func NewUserCredentialsService(repo UserCredentialsRepository, log *slog.Logger) UserCredentialsService {
	return &userCredentialsServiceImpl{
		repository: repo,
		log: log,
	}
}

func (s *userCredentialsServiceImpl) CreateUserCredentials(uc UserCredentials) (*UserCredentials, error) {
	newCredentials, err := s.repository.CreateUserCredentials(uc)
	if err != nil {
		s.log.Error("Failed creating new user credentials", "userId", uc.UserID, "error", err)
		return nil, fmt.Errorf("create user credentials failure: %w", err)
	}
	return newCredentials, nil
}

func (s *userCredentialsServiceImpl) DeleteUserCredentials(userId int64) error {
	err := s.repository.DeleteUserCredentials(userId)
	if err != nil {
		s.log.Error("Failed deleting user credentials", "userId", userId, "error", err)
		return fmt.Errorf("delete user credentials failure: %w", err)
	}
	return nil
}

func (s *userCredentialsServiceImpl) GetUserCredentials(userId int64) (*UserCredentials, error) {
	uc, err := s.repository.GetUserCredentials(userId)
	if err != nil {
		s.log.Error("Failed getting user credentials", "userId", userId, "error", err)
		return nil, fmt.Errorf("get user credentials failure: %w", err)
	}
	return uc, nil
}

func (s *userCredentialsServiceImpl) UpdateUserCredentials(uc UserCredentials) (*UserCredentials, error) {	
	updatedCredentials, err := s.repository.UpdateUserCredentials(uc)
	if err != nil {
		s.log.Error("Failed updating user credentials", "userId", uc.UserID, "error", err)
		return nil, fmt.Errorf("update user credentials failure: %w", err)
	}
	return updatedCredentials, nil
}