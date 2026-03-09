package permissions

import (
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type PermissionRepository interface {
	CountPermissions(filter repository.Filter) (int64, error)
	CreatePermission(permission Permission) (*Permission, error)
	DeletePermission(id int64) error
	FindPermissionBy(filter repository.Filter) ([]Permission, error)
	GetPermission(id int64) (*Permission, error)
	GetPermissionByCode(code string) (*Permission, error)
	UpdatePermission(permission Permission) (*Permission, error)
}

type PermissionRepositoryImpl struct {
	store database.DBStore
}

func NewPermissionRepository(store database.DBStore) PermissionRepository {
	return &PermissionRepositoryImpl{
		store: store,
	}
}

func (r *PermissionRepositoryImpl) CountPermissions(filter repository.Filter) (int64, error) {
	return r.store.Count(&Permission{}, filter)
}

func (r *PermissionRepositoryImpl) CreatePermission(permission Permission) (*Permission, error) {
	var newPerm Permission
	err := r.store.Create(&permission, permission.ToMap(), &newPerm)
	if err != nil {
		return nil, err
	}
	return &newPerm, nil
}

func (r *PermissionRepositoryImpl) DeletePermission(id int64) error {
	return r.store.Delete(&Permission{BaseModel: model.BaseModel{ ID: id }})
}

func (r *PermissionRepositoryImpl) FindPermissionBy(filter repository.Filter) ([]Permission, error) {
	var perms []Permission
	err := r.store.FindBy(&Permission{}, filter, &perms)
	return perms, err
}

func (r *PermissionRepositoryImpl) GetPermission(id int64) (*Permission, error) {
	var perm Permission
	err := r.store.Get(&Permission{ BaseModel: model.BaseModel{ ID: id }}, &perm)
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

func (r *PermissionRepositoryImpl) GetPermissionByCode(code string) (*Permission, error) {
	var perm Permission
	err := r.store.GetBy(&Permission{}, "code", code, &perm)
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

func (r *PermissionRepositoryImpl) UpdatePermission(perm Permission) (*Permission, error) {
	var updatedPerm Permission
	err := r.store.Update(&perm, perm.ToMap(), &updatedPerm)
	if err != nil {
		return nil, err
	}
	return &updatedPerm, nil
}