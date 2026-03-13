package users

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/identity/roles"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type UserRepository interface {
	AttachRole(roleID, userID int64) error
	CountUsers(filter repository.Filter) (int64, error)
	CreateUser(user User) (*User, error)
	DetachRole(roleID, userID int64) error
	DeleteUser(userId int64) error
	FindUsersBy(filter repository.Filter) ([]User, error)
	GetRolesByUserId(userId int64) ([]roles.Role, error)
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

func (r *UserRepositoryImpl) AttachRole(roleID, userID int64) error {
	return r.store.Attach(User{}.RoleLink(), userID, roleID)
}

func (r *UserRepositoryImpl) CountUsers(filter repository.Filter) (int64, error) {
	return r.store.Count(&User{}, filter)
}

func (r *UserRepositoryImpl) CreateUser(user User) (*User, error) {
	var newUser User
	err := r.store.Create(&user, user.ToMap(), &newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (r *UserRepositoryImpl) DetachRole(roleID, userID int64) error {
	return r.store.Detach(User{}.RoleLink(), userID, roleID)
}

func (r *UserRepositoryImpl) DeleteUser(userId int64) error {
	return r.store.Delete(&User{BaseModel: model.BaseModel{ID: userId}})
}

func (r *UserRepositoryImpl) FindUsersBy(filter repository.Filter) ([]User, error) {
	var users []User
	err := r.store.FindBy(&User{}, filter, &users)
	return users, err
}

func (r *UserRepositoryImpl) GetRolesByUserId(userId int64) ([]roles.Role, error) {
	var items []roles.Role
	builder := sq.Select("r.id as id, r.name as name, r.description as description, r.created_at as created_at, r.updated_at as updated_at, r.deleted_at as deleted_at").
		From("roles r").
		Join("user_roles ur ON r.id = ur.role_id").
		Where(sq.Eq{"ur.user_id": userId}).
		PlaceholderFormat(sq.Question)		
	rows, err := r.store.GetClient().QueryMany(builder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r roles.Role
		if err := rows.Scan(&r.ID, &r.Name, &r.Description, &r.CreatedAt, &r.UpdatedAt, &r.DeletedAt); err != nil {
			return nil, err
		}
		items = append(items, r)

	}
	return items, nil

}

func (r *UserRepositoryImpl) GetUser(userId int64) (*User, error) {
	var user User
	err := r.store.Get(&User{BaseModel: model.BaseModel{ID: userId}}, &user)
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