package middleware_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keithyw/pitch-in/pkg/jwt"
	"github.com/keithyw/pitch-in/pkg/middleware"
)

func TestAuthMiddleware(t *testing.T) {
    logger := slog.New(slog.NewTextHandler(io.Discard, nil))
    svc := jwt.NewJWTService("secret", 1, logger)
    mw := middleware.AuthMiddleware(svc)

    // A simple handler that returns 200 OK if the middleware lets it through
    nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    tests := []struct {
        name       string
        authHeader string
        wantStatus int
    }{
        {
            name:       "Missing Header",
            authHeader: "",
            wantStatus: http.StatusUnauthorized,
        },
		{
            name:       "Invalid Format - No Bearer",
            authHeader: "Token only-string-here",
            wantStatus: http.StatusUnauthorized, // Hits: if len(parts) != 2 || parts[0] != "Bearer"
        },
        {
            name:       "Invalid Format - Basic Auth",
            authHeader: "Basic dXNlcjpwYXNz",
            wantStatus: http.StatusUnauthorized, // Hits: same block as above
        },
        {
            name:       "Invalid Token - Wrong Secret",
            authHeader: "Bearer " + func() string {
                badSvc := jwt.NewJWTService("wrong-secret", 1, logger)
                tkn, _ := badSvc.CreateJWT(1, "user")
                return tkn
            }(),
            wantStatus: http.StatusUnauthorized, // Hits: if err != nil (after ParseJWT)
        },
        {
            name:       "Valid Token",
            authHeader: "Bearer " + func() string {
                tkn, _ := svc.CreateJWT(1, "user")
                return tkn
            }(),
            wantStatus: http.StatusOK,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/", nil)
            if tt.authHeader != "" {
                req.Header.Set("Authorization", tt.authHeader)
            }
            
            rr := httptest.NewRecorder()
            mw(nextHandler).ServeHTTP(rr, req)

            if rr.Code != tt.wantStatus {
                t.Errorf("Expected status %d, got %d", tt.wantStatus, rr.Code)
            }
        })
    }
}
