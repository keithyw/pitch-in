package credentials

import "github.com/keithyw/pitch-in/internal/database"

type UserCredentialsRepository interface {
	CreateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error)
	DeleteUserCredentials(userId int64) error
	GetUserCredentials(userId int64) (*UserCredentials, error)
	UpdateUserCredentials(userCredentials UserCredentials) (*UserCredentials, error)
}

type userCredentialsRepositoryImpl struct {
	store database.DBStore
}

func NewUserCredentialsRepository(store database.DBStore) UserCredentialsRepository {
	return &userCredentialsRepositoryImpl{
		store: store,
	}
}

func (r *userCredentialsRepositoryImpl) CreateUserCredentials(uc UserCredentials) (*UserCredentials, error) {
	var newCredentials UserCredentials
	err := r.store.Create(&uc, uc.ToMap(), &newCredentials)
	if err != nil {
		return nil, err
	}
	return &newCredentials, nil
}

func (r *userCredentialsRepositoryImpl) DeleteUserCredentials(userId int64) error {
	return r.store.Delete(&UserCredentials{UserID: userId})
}

func (r *userCredentialsRepositoryImpl) GetUserCredentials(userId int64) (*UserCredentials, error) {
	var uc UserCredentials
	err := r.store.Get(&UserCredentials{UserID: userId}, &uc)
	if err != nil {
		return nil, err
	}
	return &uc, nil
}

func (r *userCredentialsRepositoryImpl) UpdateUserCredentials(uc UserCredentials) (*UserCredentials, error) {
	var updatedCredentials UserCredentials
	err := r.store.Update(&uc, uc.ToMap(), &updatedCredentials)
	if err != nil {
		return nil, err
	}
	return &updatedCredentials, nil
}