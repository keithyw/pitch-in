package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/keithyw/pitch-in/pkg/middleware"
)

type TestPayload struct {
	Email string `validate:"required,email"`
}

func TestDecodeAndValidate(t *testing.T) {
	// Dummy handler to execute if validation passes
	finalHandler := func(w http.ResponseWriter, r *http.Request, payload TestPayload) {
		w.WriteHeader(http.StatusOK)
	}

	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{
			name:       "Success - Valid JSON and Data",
			body:       `{"email": "test@example.com"}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Failure - Malformed JSON",
			body:       `{"email": "test@example.com"`, // Missing brace
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Failure - Validation Error",
			body:       `{"email": "not-an-email"}`,
			wantStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := middleware.DecodeAndValidate(finalHandler)
			req := httptest.NewRequest("POST", "/", strings.NewReader(tt.body))
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("%s: expected status %d, got %d", tt.name, tt.wantStatus, rr.Code)
			}
		})
	}
}
