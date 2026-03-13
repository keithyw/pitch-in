package roles

import (
	"fmt"
	"log/slog"

	"github.com/keithyw/pitch-in/internal/domains/identity/permissions"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type RoleService interface {
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

type roleServiceImpl struct {
	repository RoleRepository
	log *slog.Logger
}

func NewRoleService(repo RoleRepository, log *slog.Logger) RoleService {
	return &roleServiceImpl{
		repository: repo,
		log: log,
	}
}

func (s *roleServiceImpl) AttachPermission(permissionID, roleID int64) error {
	err := s.repository.AttachPermission(permissionID, roleID)
	if err != nil {
		s.log.Error("Failed attaching permission", "permissionID", permissionID, "roleID", roleID, "error", err)
		return fmt.Errorf("attach permission failure: %w", err)
	}
	return nil
}

func (s *roleServiceImpl) CountRoles(filter repository.Filter) (int64, error) {
	cnt, err := s.repository.CountRoles(filter)
	if err != nil {
		s.log.Error("Failed getting role count", "error", err)
		return 0, fmt.Errorf("role count failure: %w", err)
	}
	return cnt, nil
}

func (s *roleServiceImpl) CreateRole(role Role) (*Role, error) {
	newRole, err := s.repository.CreateRole(role)
	if err != nil {
		s.log.Error("Failed creating role", "name", role.Name, "error", err)
		return nil, fmt.Errorf("create role error: %w", err)
	}
	return newRole, nil
}

func (s *roleServiceImpl) DeleteRole(id int64) error {
	err := s.repository.DeleteRole(id)
	if err != nil {
		s.log.Error("Failed deleting role", "id", id, "error", err)
		return fmt.Errorf("delete role error: %w", err)
	}
	return nil
}

func (s *roleServiceImpl) DetachPermission(permissionID, roleID int64) error {
	err := s.repository.DetachPermission(permissionID, roleID)
	if err != nil {
		s.log.Error("Failed detaching permission", "permissionID", permissionID, "roleID", roleID, "error", err)
		return fmt.Errorf("detach permission error: %w", err)
	}
	return nil
}

func (s *roleServiceImpl) FindRoleBy(filter repository.Filter) ([]Role, error) {
	roles, err := s.repository.FindRoleBy(filter)
	if err != nil {
		s.log.Error("Failed finding roles", "error", err)
		return nil, fmt.Errorf("Find roles by error: %w", err)
	}
	return roles, nil
}

func (s *roleServiceImpl) GetRole(id int64) (*Role, error) {
	role, err := s.repository.GetRole(id)
	if err != nil {
		s.log.Error("Failed getting role", "id", id, "error", err)
		return nil, fmt.Errorf("get role error: %w", err)
	}

	perms, err := s.GetPermissionsByRoleId(id)
	if err != nil {
		return role, nil
	}

	role.Permissions = perms
	return role, nil
}

func (s *roleServiceImpl) GetPermissionsByRoleId(id int64) ([]permissions.Permission, error) {
	perms, err := s.repository.GetPermissionsByRoleId(id)
	if err != nil {
		s.log.Error("Failing getting permissions by role", "id", id, "error", err)
		return nil, fmt.Errorf("get permissioms by role error: %w", err)
	}
	return perms, nil
}

func (s *roleServiceImpl) UpdateRole(role Role) (*Role, error) {
	updatedRole, err := s.repository.UpdateRole(role)
	if err != nil {
		s.log.Error("Failed updating role", "id", role.ID, "error", err)
		return nil, fmt.Errorf("Update role error: %w", err)
	}
	return updatedRole, nil
}