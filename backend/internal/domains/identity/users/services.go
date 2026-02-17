package users

import (
	"fmt"
	"log/slog"

	"github.com/keithyw/pitch-in/pkg/repository"
)

type UserService interface {
	CreateUser(user User) (*User, error)
	DeleteUser(userId int64) error
	FindUserBy(filter repository.Filter) ([]User, error)
	GetUser(userId int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user User) (*User, error)
}

type userServiceImpl struct {
	repository UserRepository
	log *slog.Logger
}

func NewUserService(repo UserRepository, log *slog.Logger) UserService {
	return &userServiceImpl{
		repository: repo,
		log: log,
	}
}

func (s *userServiceImpl) CreateUser(user User) (*User, error) {
	newUser, err := s.repository.CreateUser(user)
	if err != nil {
		s.log.Error("Failed creating new user", "username", user.Username, "error", err)
		return nil, fmt.Errorf("create user failure: %w", err)
	}
	return newUser, nil
}

func (s *userServiceImpl) DeleteUser(userId int64) error {
	err := s.repository.DeleteUser(userId)
	if err != nil {
		s.log.Error("Failed deleting user", "userId", userId, "error", err)
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

func (s *userServiceImpl) FindUserBy(filter repository.Filter) ([]User, error) {
	users, err := s.repository.FindUsersBy(filter)
	if err != nil {
		s.log.Error("Failed finding users", "error", err)
		return nil, fmt.Errorf("Find users by error: %w", err)
	}
	return users, nil
}

func (s *userServiceImpl) GetUser(userId int64) (*User, error) {
	user, err := s.repository.GetUser(userId)
	if err != nil {
		s.log.Error("Failed getting user", "userId", userId, "error", err)
		return nil, fmt.Errorf("Get user error: %w", err)
	}
	return user, nil
}

func (s *userServiceImpl) GetUserByEmail(email string) (*User, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		s.log.Error("Failed getting user by email", "email", email, "error", err)
		return nil, fmt.Errorf("Get user by email error: %w", err)
	}
	return user, nil
}

func (s *userServiceImpl) UpdateUser(user User) (*User, error) {
	updatedUser, err := s.repository.UpdateUser(user)
	if err != nil {
		s.log.Error("Failed updating user", "userId", user.ID, "error", err)
		return nil, fmt.Errorf("Update user error: %w", err)
	}
	return updatedUser, nil
}