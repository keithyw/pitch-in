package users

import (
	"time"

	"github.com/keithyw/pitch-in/pkg/model"
)

type UserFields struct {
	Username *string `json:"username,omitempty" db:"username" validate:"omitempty,min=3,max=255"`
	Email *string `json:"email,omitempty" db:"email" validate:"omitempty,email"`
	FirstName *string `json:"first_name,omitempty" db:"first_name" validate:"omitempty,min=2,max=255"`
	LastName *string `json:"last_name,omitempty" db:"last_name" validate:"omitempty,min=2,max=255"`
	IsActive *bool `json:"is_active" db:"is_active" validate:"boolean"`
}
type User struct {
	ID int64 `json:"id" db:"id"`
	UserFields
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type PatchUserRequest struct {
	UserFields
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Columns() []string{
	return []string{"id", "username", "email", "first_name", "last_name", "is_active", "created_at", "updated_at", "deleted_at"}
}

func (u *User) IsAutoIncrementKey() bool {
	return true
}

func (u *User) PrimaryKey() (string, interface{}) {
	return "id", u.ID
}

func (u *User) SetID(id int64) {
	u.ID = id
}

func (u *User) ToMap() map[string]interface{} {
	fields := map[string]interface{}{
		"username": u.Username,
		"email": u.Email,
		"first_name": u.FirstName,
		"last_name": u.LastName,
		"is_active": u.IsActive,
	}
	m := model.MapValues(fields)
	return m
}

func (p *PatchUserRequest) ToModel(id int64) *User {
	return &User{
		ID: id,
		UserFields: p.UserFields,
	}
}