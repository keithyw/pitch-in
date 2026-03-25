package items

import (
	"github.com/keithyw/pitch-in/internal/domains/taxonomy/tags"
	"github.com/keithyw/pitch-in/pkg/model"
)

type ItemFields struct {
	Name *string `json:"name,omitempty" db:"name" validate:"omitempty,max=255"`
	Slug *string `json:"slug,omitempty" db:"slug" validate:"omitempty,max=255"`
	Description *string `json:"description,omitempty" db:"description" validate:"omitempty"`
}

type Item struct{
	model.BaseModel
	ItemFields
	Tags []tags.Tag `json:"tags"`
}

type PatchItemRequest struct {
	ItemFields
}

func (m *Item) TableName() string {
	return "items"
}

func (m *Item) Columns() []string {
	return []string{"id", "name", "slug", "description", "created_at", "updated_at", "deleted_at"}
}

func (m *Item) ToMap() map[string]interface{} {
	fields := map[string]interface{} {
		"name": m.Name,
		"slug": m.Slug,
		"description": m.Description,
	}
	return model.MapValues(fields)
}

func (m *PatchItemRequest) ToModel(id int64) *Item {
	return &Item{
		BaseModel: model.BaseModel{
			ID: id,
		},
		ItemFields: m.ItemFields,
	}
}