package users

import "time"

type User struct {
	ID int64 `json:"id" db:"id"`
	Username string `json:"username" db:"username" validate:"required,min=3,max=255"`
	Email string `json:"email" db:"email" validate:"required,email"`
	FirstName string `json:"first_name,omitempty" db:"first_name" validate:"min=2,max=255"`
	LastName string `json:"last_name,omitempty" db:"last_name" validate:"min=2,max=255"`
	IsActive bool `json:"is_active" db:"is_active" validate:"boolean"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) Columns() []string{
	return []string{"id", "username", "email", "first_name", "last_name", "is_active", "created_at", "updated_at", "deleted_at"}
}

func (u *User) PrimaryKey() (string, interface{}) {
	return "id", u.ID
}

func (u *User) SetID(id int64) {
	u.ID = id
}

func (u *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"username": u.Username,
		"email": u.Email,
		"first_name": u.FirstName,
		"last_name": u.LastName,
		"is_active": u.IsActive,
	}
}
