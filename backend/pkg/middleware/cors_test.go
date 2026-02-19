package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keithyw/pitch-in/pkg/middleware"
)

func TestCors(t *testing.T) {
	mw := middleware.Cors
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{"Standard GET Request", http.MethodGet, http.StatusOK},
		{"Preflight OPTIONS Request", http.MethodOptions, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/", nil)
			rr := httptest.NewRecorder()

			mw(nextHandler).ServeHTTP(rr, req)

			// Check Headers
			if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
				t.Error("CORS header not set")
			}

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}
