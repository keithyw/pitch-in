package permissions

import (
	"fmt"
	"log/slog"

	"github.com/keithyw/pitch-in/pkg/repository"
)

type PermissionService interface {
	CountPermissions(filter repository.Filter) (int64, error)
	CreatePermission(permission Permission) (*Permission, error)
	DeletePermission(id int64) error
	FindPermissionBy(filter repository.Filter) ([]Permission, error)
	GetPermission(id int64) (*Permission, error)
	GetPermissionByCode(code string) (*Permission, error)
	UpdatePermission(permission Permission) (*Permission, error)
}

type permissionServiceImpl struct {
	repository PermissionRepository
	log *slog.Logger
}

func NewPermissionService(repo PermissionRepository, log *slog.Logger) PermissionService {
	return &permissionServiceImpl{
		repository: repo,
		log: log,
	}
}

func (s *permissionServiceImpl) CountPermissions(filter repository.Filter) (int64, error) {
	cnt, err := s.repository.CountPermissions(filter)
	if err != nil {
		s.log.Error("Failed getting permission count", "error", err)
		return 0, fmt.Errorf("permission count failure: %w", err)
	}
	return cnt, nil
}

func (s *permissionServiceImpl) CreatePermission(permission Permission) (*Permission, error) {
	newPerm, err := s.repository.CreatePermission(permission)
	if err != nil {
		s.log.Error("Failed creating new permission", "code", permission.Code, "error", err)
		return nil, fmt.Errorf("create permission error: %w", err)
	}
	return newPerm, nil
}

func (s *permissionServiceImpl) DeletePermission(id int64) error {
	err := s.repository.DeletePermission(id)
	if err != nil {
		s.log.Error("Failed deleting permission", "id", id, "error", err)
		return fmt.Errorf("delete permission error: %w", err)
	}
	return nil
}

func (s *permissionServiceImpl) FindPermissionBy(filter repository.Filter) ([]Permission, error) {
	perms, err := s.repository.FindPermissionBy(filter)
	if err != nil {
		s.log.Error("Failed finding permission", "error", err)
		return nil, fmt.Errorf("Find permissions by error: %w", err)
	}
	return perms, nil
}

func (s *permissionServiceImpl) GetPermission(id int64) (*Permission, error) {
	perm, err := s.repository.GetPermission(id)
	if err != nil {
		s.log.Error("Failed getting permission", "id", id, "error", err)
		return nil, fmt.Errorf("Get permission error: %w", err)
	}
	return perm, nil
}

func (s *permissionServiceImpl) GetPermissionByCode(code string) (*Permission, error) {
	perm, err := s.repository.GetPermissionByCode(code)
	if err != nil {
		s.log.Error("Failed getting permission by code", "code", code, "error", err)
		return nil, fmt.Errorf("Get permission by code error: %w", err)
	}
	return perm, nil
}

func (s *permissionServiceImpl) UpdatePermission(permission Permission) (*Permission, error) {
	updatedPerm, err := s.repository.UpdatePermission(permission)
	if err != nil {
		s.log.Error("Failed updating permission", "id", permission.ID, "error", err)
		return nil, fmt.Errorf("Update permission error: %w", err)
	}
	return updatedPerm, nil
}