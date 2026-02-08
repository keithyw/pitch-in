package repository

type Filter struct {
	Fields map[string]interface{}
	Operators map[string]string
	Limit int
	Offset int
	Order string
	Sort string
}