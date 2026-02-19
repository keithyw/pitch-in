package users

import (
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user User) (*User, error) {
	args := m.Called(user)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) DeleteUser(userId int64) error {
	return m.Called(userId).Error(0)
}

func (m *MockUserService) FindUserBy(filter repository.Filter) ([]User, error) {
	args := m.Called(filter)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockUserService) GetUser(userId int64) (*User, error) {
	args := m.Called(userId)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) GetUserByEmail(email string) (*User, error) {
	args := m.Called(email)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user User) (*User, error) {
	args := m.Called(user)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}