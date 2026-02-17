package credentials

import (
	"time"

	"github.com/keithyw/pitch-in/pkg/model"
)

type UserCredentials struct {
	UserID int64 `json:"user_id" db:"user_id"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	LastLogin *time.Time `json:"last_login,omitempty" db:"last_login"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (uc *UserCredentials) TableName() string {
	return "user_credentials"
}

func (uc *UserCredentials) Columns() []string {
	return []string{"user_id", "password_hash", "last_login", "created_at", "updated_at", "deleted_at"}
}

func (uc *UserCredentials) IsAutoIncrementKey() bool {
	return false
}

func (uc *UserCredentials) PrimaryKey() (string, interface{}) {
	return "user_id", uc.UserID
}

func (uc *UserCredentials) SetID(id int64) {
	uc.UserID = id
}

func (uc *UserCredentials) ToMap() map[string]interface{} {
	fields := map[string]interface{}{
		"user_id": uc.UserID,
		"password_hash": uc.PasswordHash,
	}
	return model.MapValues(fields)
}