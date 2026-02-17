package model

import "reflect"

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