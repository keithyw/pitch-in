package users

import (
	"fmt"
	"io"
	"log/slog"
	"testing"

	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateUser(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	
	
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		svc := NewUserService(mockRepo, logger)
		username := "newuser"
		inputUser := User{UserFields: UserFields{Username: &username}}
		mockRepo.On("CreateUser", mock.Anything).Return(&User{ID: 1, UserFields: UserFields{Username: &username}}, nil)
		result, err := svc.CreateUser(inputUser)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), result.ID)
		assert.Equal(t, username, *result.Username)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		svc := NewUserService(mockRepo, logger)
        mockRepo.On("CreateUser", mock.Anything).Return(nil, fmt.Errorf("db fail")).Once()
        result, err := svc.CreateUser(User{})
        assert.Nil(t, result)
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "create user failure")
        mockRepo.AssertExpectations(t)
    })
}

func TestUserService_DeleteUser(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		svc := NewUserService(mockRepo, logger)
		mockRepo.On("DeleteUser", mock.Anything).Return(nil)
		err := svc.DeleteUser(1)
		assert.NoError(t, err)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		svc := NewUserService(mockRepo, logger)
        mockRepo.On("DeleteUser", int64(1)).Return(fmt.Errorf("not found")).Once()
        err := svc.DeleteUser(1)
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "delete user error")
        mockRepo.AssertExpectations(t)
    })
}

func TestUserService_FindUserBy(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo, logger)

	t.Run("Success - Multiple Users", func(t *testing.T) {
		filter := repository.Filter{
			Limit: 10,
			Fields: map[string]interface{}{"is_active": true},
		}
		
		expectedUsers := []User{
			{ID: 1},
			{ID: 2},
		}

		mockRepo.On("FindUsersBy", filter).Return(expectedUsers, nil).Once()

		results, err := svc.FindUserBy(filter)

		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, int64(1), results[0].ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		filter := repository.Filter{Limit: 5}
		mockRepo.On("FindUsersBy", filter).Return(nil, fmt.Errorf("db connection lost")).Once()

		results, err := svc.FindUserBy(filter)

		assert.Nil(t, results)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Find users by error")
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUser(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo, logger)

	t.Run("Success", func(t *testing.T) {
		expectedUser := &User{ID: 1}
		mockRepo.On("GetUser", int64(1)).Return(expectedUser, nil).Once()

		user, err := svc.GetUser(1)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), user.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("GetUser", int64(99)).Return(nil, fmt.Errorf("db error")).Once()

		user, err := svc.GetUser(99)

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Get user error")
		mockRepo.AssertExpectations(t)
	})
}

func TestUserService_GetUserByEmail(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	mockRepo := new(MockUserRepository)
	svc := NewUserService(mockRepo, logger)
	email := "test@test.com"

	t.Run("Success", func(t *testing.T) {
		expectedUser := &User{ID: 1, UserFields: UserFields{ Email: &email }}
		mockRepo.On("GetUserByEmail", email).Return(expectedUser, nil).Once()

		user, err := svc.GetUserByEmail(email)

		fmt.Printf("user %s", *expectedUser.Email)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), user.ID)
		assert.Equal(t, email, *user.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.On("GetUserByEmail", email).Return(nil, fmt.Errorf("db error")).Once()

		user, err := svc.GetUserByEmail(email)

		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Get user by email error: db error")
		mockRepo.AssertExpectations(t)
	})

}

func TestUserService_UpdateUser(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	t.Run("Success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		svc := NewUserService(mockRepo, logger)
		firstName := "updated"
		item := User{ID: 1, UserFields: UserFields{FirstName: &firstName, }}
		mockRepo.On("UpdateUser", mock.Anything).Return(&User{ID: 1, UserFields: UserFields{FirstName: &firstName}}, nil)
		result, err := svc.UpdateUser(item)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), result.ID)
		assert.Equal(t, item.FirstName, result.FirstName)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		svc := NewUserService(mockRepo, logger)
        mockRepo.On("UpdateUser", mock.Anything).Return(nil, fmt.Errorf("lock error")).Once()
        result, err := svc.UpdateUser(User{ID: 1})
        assert.Nil(t, result)
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "Update user error")
        mockRepo.AssertExpectations(t)
    })	
}