package tags

import "github.com/keithyw/pitch-in/pkg/model"

type TagFields struct {
	Tag *string `json:"tag,omitempty" db:"tag" validate:"omitempty,max=255"`
	Slug *string `json:"slug,omitempty" db:"slug" validate:"omitempty,max=255"`
}

type Tag struct {
	model.BaseModel
	TagFields
}

type PatchTagRequest struct {
	TagFields
}

func (t *Tag) TableName() string {
	return "tags"
}

func (t *Tag) Columns() []string {
	return []string{"id", "tag", "slug", "created_at", "updated_at", "deleted_at"}
}

func (t *Tag) ToMap() map[string]interface{} {
	fields := map[string]interface{} {
		"tag": t.Tag,
		"slug": t.Slug,
	}
	return model.MapValues(fields)
}

func (t *PatchTagRequest) ToModel(id int64) *Tag {
	return &Tag{
		BaseModel: model.BaseModel{
			ID: id,
		},
		TagFields: t.TagFields,
	}
}