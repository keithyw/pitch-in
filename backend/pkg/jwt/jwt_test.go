package jwt_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/keithyw/pitch-in/pkg/jwt"
)

func TestJWTService_ParseJWT(t *testing.T) {
	secret := "secret"
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	svc := jwt.NewJWTService(secret, 1, logger)

	validToken, _ := svc.CreateJWT(1, "testuser")
	expiredSvc := jwt.NewJWTService(secret, -1, logger)
	expiredToken, _ := expiredSvc.CreateJWT(1, "testuser")

	tests := []struct {
		name string
		tokenString string
		wantErr bool
		errContains string
	}{
		{
			name:        "Success - Valid Token",
			tokenString: validToken,
			wantErr:     false,
		},
		{
			name:        "Failure - Expired Token",
			tokenString: expiredToken,
			wantErr:     true,
			errContains: "token is expired",
		},
		{
			name:        "Failure - Malformed Token",
			tokenString: "not.a.real.token",
			wantErr:     true,
			errContains: "token is malformed",
		},
		{
			name:        "Failure - Wrong Secret",
			tokenString: func() string {
				otherSvc := jwt.NewJWTService("wrong-secret", 1, logger)
				tkn, _ := otherSvc.CreateJWT(1, "user")
				return tkn
			}(),
			wantErr:     true,
			errContains: "signature is invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims, err := svc.ParseJWT(tt.tokenString)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && claims.Username != "testuser" {
				t.Errorf("ParseJWT() got username = %v, want %v", claims.Username, "testuser")
			}
		})
	}
	
}