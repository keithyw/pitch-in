package users

import (
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type UserRepository interface {
	CreateUser(user User) (*User, error)
	DeleteUser(userId int64) error
	FindUsersBy(filter repository.Filter) ([]User, error)
	GetUser(userId int64) (*User, error)
	GetUserByEmail (email string) (*User, error)
	UpdateUser(user User) (*User, error)
}

type UserRepositoryImpl struct {
	store database.DBStore
}

func NewUserRepository(store database.DBStore) UserRepository {
	return &UserRepositoryImpl{
		store: store,
	}
}

func (r *UserRepositoryImpl) CreateUser(user User) (*User, error) {
	var newUser User
	err := r.store.Create(&user, user.ToMap(), &newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (r *UserRepositoryImpl) DeleteUser(userId int64) error {
	return r.store.Delete(&User{ID: userId})
}

func (r *UserRepositoryImpl) FindUsersBy(filter repository.Filter) ([]User, error) {
	var users []User
	err := r.store.FindBy(&User{}, filter, &users)
	return users, err
}

func (r *UserRepositoryImpl) GetUser(userId int64) (*User, error) {
	var user User
	err := r.store.Get(&User{ID: userId}, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.store.GetBy(&User{}, "email", email, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user User) (*User, error) {
	var updatedUser User
	err := r.store.Update(&user, user.ToMap(), &updatedUser)
	if err != nil { 
		return nil, err
	}
	return &updatedUser, nil
}