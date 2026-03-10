package roles

import (
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type RoleRepository interface {
	CountRoles(filter repository.Filter) (int64, error)
	CreateRole(role Role) (*Role, error)
	DeleteRole(id int64) error
	FindRoleBy(filter repository.Filter) ([]Role, error)
	GetRole(id int64) (*Role, error)
	UpdateRole(role Role) (*Role, error)
}

type RoleRepositoryImpl struct {
	store database.DBStore
}

func NewRoleRepository(store database.DBStore) RoleRepository {
	return &RoleRepositoryImpl{
		store: store,
	}
}

func (r *RoleRepositoryImpl) CountRoles(filter repository.Filter) (int64, error) {
	return r.store.Count(&Role{}, filter)
}

func (r *RoleRepositoryImpl) CreateRole(role Role) (*Role, error) {
	var newRole Role
	err := r.store.Create(&role, role.ToMap(), &newRole)
	if err != nil {
		return nil, err
	}
	return &newRole, nil
}

func (r *RoleRepositoryImpl) DeleteRole(id int64) error {
	return r.store.Delete(&Role{BaseModel: model.BaseModel{ ID: id }})
}

func (r *RoleRepositoryImpl) FindRoleBy(filter repository.Filter) ([]Role, error) {
	var roles []Role
	err := r.store.FindBy(&Role{}, filter, &roles)
	return roles, err
}

func (r *RoleRepositoryImpl) GetRole(id int64) (*Role, error) {
	var role Role
	err := r.store.Get(&Role{ BaseModel: model.BaseModel{ ID: id }}, &role)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) UpdateRole(role Role) (*Role, error) {
	var updatedRole Role
	err := r.store.Update(&role, role.ToMap(), &updatedRole)
	if err != nil {
		return nil, err
	}
	return &updatedRole, nil
}