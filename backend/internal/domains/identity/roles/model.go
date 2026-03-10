package roles

import "github.com/keithyw/pitch-in/pkg/model"

type RoleFields struct {
	Name *string `json:"name,omitempty" db:"name" validate:"omitempty,min=3,max=255"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty"`	
}

type Role struct {
	model.BaseModel
	RoleFields
}

type PatchRoleRequest struct {
	RoleFields
}

func (r *Role) TableName() string {
	return "roles"
}

func (r *Role) Columns() []string{
	return []string{"id", "name", "description", "created_at", "updated_at", "deleted_at"}
}

func (r *Role) ToMap() map[string]interface{} {
	fields := map[string]interface{} {
		"name": r.Name,
		"description": r.Description,
	}
	return model.MapValues(fields)
}

func (p *PatchRoleRequest) ToModel(id int64) *Role {
	return &Role{
		BaseModel: model.BaseModel{
			ID: id,
		},
		RoleFields: p.RoleFields,
	}
}