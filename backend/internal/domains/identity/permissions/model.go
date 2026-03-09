package permissions

import "github.com/keithyw/pitch-in/pkg/model"

type PermissionFields struct {
	Code *string `json:"code,omitempty" db:"code" validate:"omitempty,min=3,max=150"`
	DisplayName *string `json:"display_name,omitempty" db:"display_name" validate:"omitempty,max=255"`
	Path *string `json:"path,omitempty" db:"path" validate:"omitempty,max=255"`
	Method *string `json:"method,omitempty" db:"method" validate:"omitempty,max=255"`
}

type Permission struct {
	model.BaseModel
	PermissionFields
}

type PatchPermissionRequest struct {
	PermissionFields
}

func (p *Permission) TableName() string {
	return "permissions"
}

func (p *Permission) Columns() []string{
	return []string{"id", "code", "display_name", "path", "method", "created_at", "updated_at", "deleted_at"}
}

func (p *Permission) ToMap() map[string]interface{} {
	fields := map[string]interface{}{
		"code": p.Code,
		"display_name": p.DisplayName,
		"path": p.Path,
		"method": p.Method,
	}
	m := model.MapValues(fields)
	return m
}

func (p *PatchPermissionRequest) ToModel(id int64) *Permission {
	return &Permission{
		BaseModel: model.BaseModel{
			ID: id,
		},
		PermissionFields: p.PermissionFields,
	}
}