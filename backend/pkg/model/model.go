package model

type Model interface {
	TableName() string
	Columns() []string
	PrimaryKey() (column string, value interface{})
	SetID(id int64)
	ToMap() map[string]interface{}
}