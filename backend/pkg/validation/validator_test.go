package validation_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/keithyw/pitch-in/pkg/validation"
)

func TestCheckPassword(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("password_complex", validation.CheckPassword)

	type User struct {
		Password string `validate:"password_complex"`
	}

	tests := []struct {
		name     string
		password string
		valid    bool
	}{
		{"All requirements met", "Password123!", true},
		{"Missing Uppercase", "password123!", false},
		{"Missing Special", "Password123", false},
		{"Missing Digit", "Password!", false},
		{"Too Simple", "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := User{Password: tt.password}
			err := validate.Struct(user)
			if (err == nil) != tt.valid {
				t.Errorf("Password %q: expected validity %v, got %v", tt.password, tt.valid, err == nil)
			}
		})
	}
}