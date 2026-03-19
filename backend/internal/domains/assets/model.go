package assets

import "github.com/keithyw/pitch-in/pkg/model"

type AssetFields struct {
	ObjectKey *string `json:"object_key,omitempty" db:"object_key" validate:"omitempty,max=500"`
	MimeType *string `json:"mime_type,omitempty" db:"mime_type" validate:"omitempty,max=100"`
	Width *int `json:"width,omitempty" db:"width" validate:"omitempty"`
	Height *int `json:"height,omitempty" db:"height" validate:"omitempty"`
	SizeBytes *int64 `json:"size_bytes,omitempty" db:"size_bytes" validate:"omitempty"`
}

type Asset struct {
	model.BaseModel
	AssetFields
	URL *string `json:"url,omitempty" validate:"omitempty"`
}

func (a *Asset) TableName() string {
	return "assets"
}

func (a *Asset) Columns() []string {
	return []string{"id", "object_key", "mime_type", "width", "height", "size_bytes", "created_at", "updated_at", "deleted_at"}
}

func (a *Asset) ToMap() map[string]interface{} {
	fields := map[string]interface{}{
		"object_key": a.ObjectKey,
		"mime_type": a.MimeType,
		"width": a.Width,
		"height": a.Height,
		"size_bytes": a.SizeBytes,
	}
	return model.MapValues(fields)
}