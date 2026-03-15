package roles

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/identity/permissions"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type RoleRepository interface {
	AttachPermission(permissionID, roleID int64) error
	CountRoles(filter repository.Filter) (int64, error)
	CreateRole(role Role) (*Role, error)
	DeleteRole(id int64) error
	DetachPermission(permissionID, roleID int64) error
	FindRoleBy(filter repository.Filter) ([]Role, error)
	GetRole(id int64) (*Role, error)
	GetPermissionsByRoleId(id int64) ([]permissions.Permission, error)
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

func (r *RoleRepositoryImpl) AttachPermission(permissionID, roleID int64) error {
	return r.store.Attach(Role{}.PermissionLink(), roleID, permissionID)
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

func (r *RoleRepositoryImpl) DetachPermission(permissionID, roleID int64) error {
	return r.store.Detach(Role{}.PermissionLink(), roleID, permissionID)
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

func (r *RoleRepositoryImpl) GetPermissionsByRoleId(id int64) ([]permissions.Permission, error) {
	var perms []permissions.Permission
	builder := sq.Select("p.id as id, p.code as code, p.display_name as display_name, p.path as path, p.method as method, p.created_at as created_at, p.updated_at as updated_at, p.deleted_at as deleted_at").
		From("permissions p").
		Join("role_permissions rp ON p.id = rp.permission_id").
		Where(sq.Eq{"rp.role_id": id}).
		PlaceholderFormat(sq.Question)		
	rows, err := r.store.GetClient().QueryMany(builder)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p permissions.Permission
		if err := rows.Scan(&p.ID, &p.Code, &p.DisplayName, &p.Path, &p.Method, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
			return nil, err
		}
		perms = append(perms, p)
	}
	return perms, nil
}

func (r *RoleRepositoryImpl) UpdateRole(role Role) (*Role, error) {
	var updatedRole Role
	err := r.store.Update(&role, role.ToMap(), &updatedRole)
	if err != nil {
		return nil, err
	}
	return &updatedRole, nil
}