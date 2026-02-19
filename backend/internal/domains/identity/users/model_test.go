package users

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUser_ModelInterface(t *testing.T) {
	username := "keithyw"
	email := "test@example.com"
	user := &User{
		ID: 100,
		UserFields: UserFields{
			Username: &username,
			Email:    &email,
		},
	}

	t.Run("Metadata", func(t *testing.T) {
		assert.Equal(t, "users", user.TableName())
		assert.True(t, user.IsAutoIncrementKey())
		assert.Contains(t, user.Columns(), "username")
	})

	t.Run("Primary Key Logic", func(t *testing.T) {
		col, val := user.PrimaryKey()
		assert.Equal(t, "id", col)
		assert.Equal(t, int64(100), val)

		user.SetID(200)
		assert.Equal(t, int64(200), user.ID)
	})

	t.Run("ToMap Transformation", func(t *testing.T) {
		data := user.ToMap()
		
		assert.Equal(t, "keithyw", *data["username"].(*string))
		assert.Equal(t, "test@example.com", *data["email"].(*string))
		assert.Nil(t, data["first_name"])
	})
}

func TestUser_Validation(t *testing.T) {
	invalidEmail := "not-an-email"
	user := &User{
		UserFields: UserFields{
			Email: &invalidEmail,
		},
	}

	v := validator.New() 
	err := v.Struct(user)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Email")
}

func TestPatchUserRequest_ToModel(t *testing.T) {
	username := "keith_patch"
	req := PatchUserRequest{
		UserFields: UserFields{
			Username: &username,
		},
	}

	user := req.ToModel(123)
	assert.NotNil(t, user)
	assert.Equal(t, username, *user.Username)
	assert.Nil(t, user.FirstName)
}