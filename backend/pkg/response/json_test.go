package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/keithyw/pitch-in/pkg/response"
)

func TestJSON(t *testing.T) {
	tests := []struct {
		name       string
		data       interface{}
		status     int
		wantBody   string
	}{
		{
			name:       "Standard Object",
			data:       map[string]string{"message": "success"},
			status:     http.StatusOK,
			wantBody:   "{\"message\":\"success\"}\n",
		},
		{
			name:       "Nil Slice to Empty Array",
			data:       []string(nil),
			status:     http.StatusOK,
			wantBody:   "[]\n",
		},
		{
			name:       "Error JSON helper",
			data:       "error message",
			status:     http.StatusBadRequest,
			wantBody:   "{\"error\":\"error message\"}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			if tt.name == "Error JSON helper" {
				response.ErrorJSON(rr, tt.status, tt.data.(string))
			} else {
				response.JSON(rr, tt.status, tt.data)
			}

			if rr.Code != tt.status {
				t.Errorf("expected status %d, got %d", tt.status, rr.Code)
			}
			if rr.Body.String() != tt.wantBody {
				t.Errorf("expected body %q, got %q", tt.wantBody, rr.Body.String())
			}
		})
	}
}