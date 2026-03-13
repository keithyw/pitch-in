package users

import (
	"github.com/keithyw/pitch-in/internal/domains/identity/roles"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) AttachRole(roleID, userID int64) error {
	return m.Called(roleID, userID).Error(0)
}

func (m *MockUserRepository) CountUsers(filter repository.Filter) (int64, error) {
	args := m.Called(filter)
	if args.Get(0) == nil { return 0, args.Error(1) }
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user User) (*User, error) {
	args := m.Called(user)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(userId int64) error {
	return m.Called(userId).Error(0)
}

func (m *MockUserRepository) DetachRole(roleID, userID int64) error {
	return m.Called(roleID, userID).Error(0)
}

func (m *MockUserRepository) FindUsersBy(filter repository.Filter) ([]User, error) {
	args := m.Called(filter)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockUserRepository) GetRolesByUserId(userId int64) ([]roles.Role, error) {
	args := m.Called(userId)
	if args.Get(0) == nil { return nil, args.Error(1) }
	return args.Get(0).([]roles.Role), args.Error(1)
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