package repository_test

import (
	"io"
	"log/slog"
	"net/url"
	"testing"

	"github.com/keithyw/pitch-in/pkg/repository"
)

type MockModel struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"full_name"`
}

func TestParser_Parse(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	p := repository.NewParser(MockModel{}, log)

	tests := []struct {
		name    string
		params  url.Values
		wantErr bool
		check   func(*testing.T, *repository.Filter)
	}{
		{
			name: "Basic Pagination and Sort",
			params: url.Values{
				"limit":  {"10"},
				"offset": {"5"},
				"sort":   {"name.desc"},
			},
			wantErr: false,
			check: func(t *testing.T, f *repository.Filter) {
				if f.Limit != 10 || f.Offset != 5 || f.Sort != "full_name" || f.Order != "desc" {
					t.Errorf("Unexpected filter values: %+v", f)
				}
			},
		},
		{
			name: "Operator Parsing - Between",
			params: url.Values{
				"id__between": {"1|100"},
			},
			wantErr: false,
			check: func(t *testing.T, f *repository.Filter) {
				if f.Operators["id"] != "between" {
					t.Errorf("Expected operator between, got %s", f.Operators["id"])
				}
			},
		},
		{
			name: "Invalid Sort Field",
			params: url.Values{
				"sort": {"invalid_field"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter, err := p.Parse(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.check != nil {
				tt.check(t, filter)
			}
		})
	}
}