package model

import (
	"reflect"
	"time"
)


type BaseModel struct {
	ID int64 `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (b *BaseModel) PrimaryKey() (string, interface{}) {
	return "id", b.ID
}

func (b *BaseModel) IsAutoIncrementKey() bool {
	return true
}

func (b *BaseModel) SetID(id int64) {
	b.ID = id
}

type Model interface {
	TableName() string
	Columns() []string
	IsAutoIncrementKey() bool
	PrimaryKey() (column string, value interface{})
	SetID(id int64)
	ToMap() map[string]interface{}
}

func MapValues(fields map[string]interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	for key, v := range fields {
		if v != "" {
			val := reflect.ValueOf(v)
			if val.Kind() == reflect.Ptr && val.IsNil() {
				continue
			}
			m[key] = v
		}
	}
	return m
}