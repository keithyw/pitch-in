package users

import (
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user User) (*User, error) {
	args := m.Called(user)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(userId int64) error {
	return m.Called(userId).Error(0)
}

func (m *MockUserRepository) FindUsersBy(filter repository.Filter) ([]User, error) {
	args := m.Called(filter)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockUserRepository) GetUser(userId int64) (*User, error) {
	args := m.Called(userId)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*User, error) {
	args := m.Called(email)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user User) (*User, error) {
	args := m.Called(user)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}