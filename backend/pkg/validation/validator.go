package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CheckPassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
        hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
        hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
        hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
        hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
        
        return hasUpper && hasLower && hasDigit && hasSpecial
}