package database_test

import (
	"context"
	"fmt"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/keithyw/pitch-in/internal/database"
	"github.com/keithyw/pitch-in/internal/domains/identity/users"
	"github.com/keithyw/pitch-in/pkg/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockResult struct { lastId int64 }
func (m mockResult) LastInsertId() (int64, error) { return m.lastId, nil }
func (m mockResult) RowsAffected() (int64, error) { return 1, nil }

func TestDBStore_Errors(t *testing.T) {
    mockClient := new(database.MockDBClient)
    store := database.NewDBStore(context.Background(), mockClient)
    m := &users.User{ID: 1}

    // Simulate a DB failure on Delete
    mockClient.On("Exec", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("connection refused"))

    err := store.Delete(m)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "connection refused")
}
func TestMakeQueryFromFilter(t *testing.T) {
	// We can instantiate the internal struct directly for white-box testing
	mockClient := new(database.MockDBClient)
	store := database.NewDBStore(context.Background(), mockClient)
	baseQuery := sq.Select("*").From("users")

	tests := []struct {
		name     string
		filter   repository.Filter
		wantSQL  string
		wantArgs []interface{}
	}{
		{
			name: "Basic Equality",
			filter: repository.Filter{
				Fields: map[string]interface{}{"status": "active"},
			},
			wantSQL:  "SELECT * FROM users WHERE status = ?",
			wantArgs: []interface{}{"active"},
		},
		{
			name: "Greater Than and Less Than",
			filter: repository.Filter{
				Fields:    map[string]interface{}{"age": 18, "score": 100},
				Operators: map[string]string{"age": ">", "score": "<="},
			},
			wantSQL:  "SELECT * FROM users WHERE age > ? AND score <= ?",
			wantArgs: []interface{}{18, 100},
		},
		{
			name: "Full Text Search (Match Against)",
			filter: repository.Filter{
				Fields:    map[string]interface{}{"bio": "developer"},
				Operators: map[string]string{"bio": "~="},
			},
			wantSQL:  "SELECT * FROM users WHERE MATCH (bio) AGAINST (? IN BOOLEAN MODE)",
			wantArgs: []interface{}{"developer"},
		},
		{
			name: "In Operator",
			filter: repository.Filter{
				Fields:    map[string]interface{}{"id": []string{"1", "2", "3"}},
				Operators: map[string]string{"id": "in"},
			},
			wantSQL:  "SELECT * FROM users WHERE id IN (?,?,?)",
			wantArgs: []interface{}{"1", "2", "3"},
		},
		{
			name: "Between Operator",
			filter: repository.Filter{
				Fields:    map[string]interface{}{"created_at": []string{"2023-01-01", "2023-12-31"}},
				Operators: map[string]string{"created_at": "between"},
			},
			wantSQL:  "SELECT * FROM users WHERE created_at BETWEEN ? AND ?",
			wantArgs: []interface{}{"2023-01-01", "2023-12-31"},
		},
		{
			name: "Skip Empty or Nil Values",
			filter: repository.Filter{
				Fields: map[string]interface{}{"name": "", "deleted_at": nil},
			},
			wantSQL:  "SELECT * FROM users",
			wantArgs: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := store.MakeQueryFromFilter(tt.filter, baseQuery)
			sql, args, err := query.ToSql()

			assert.NoError(t, err)
			assert.Equal(t, tt.wantSQL, sql)
			assert.Equal(t, tt.wantArgs, args)
		})
	}
}

func TestDBStore_Create(t *testing.T) {
	mockClient := new(database.MockDBClient)
	store := database.NewDBStore(context.Background(), mockClient)
	
	// Define a mock model (assuming you have one for testing)
	m := &users.User{ID: 0} 
	data := map[string]interface{}{"username": "keith"}

	// Set expectation: Create should call Exec with an Insert builder
	mockClient.On("Exec", mock.Anything, mock.Anything).Return(mockResult{lastId: 10}, nil)
	// Set expectation: After create, it calls Get to refresh the object
	mockClient.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := store.Create(m, data, m)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestDBStore_Delete(t *testing.T) {
	mockClient := new(database.MockDBClient)
	store := database.NewDBStore(context.Background(), mockClient)
	
	m := &users.User{ID: 1}

	// Expect Exec to be called with a DELETE query
	mockClient.On("Exec", mock.Anything, mock.Anything).Return(mockResult{}, nil)

	err := store.Delete(m)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestDBStore_FindBy(t *testing.T) {
	mockClient := new(database.MockDBClient)
	store := database.NewDBStore(context.Background(), mockClient)
	m := &users.User{}
	
	filter := repository.Filter{
		Limit: 10,
		Fields: map[string]interface{}{"role": "admin"},
	}

	var results []users.User
	// Expect Query (SelectContext) to be called
	mockClient.On("Query", mock.Anything, mock.MatchedBy(func(q sq.Sqlizer) bool {
		sql, _, _ := q.ToSql()
		return assert.Contains(t, sql, "LIMIT 10") && assert.Contains(t, sql, "WHERE role = ?")
	}), &results).Return(nil)

	err := store.FindBy(m, filter, &results)
	assert.NoError(t, err)
}

func TestDBStore_GetBy(t *testing.T) {
	mockClient := new(database.MockDBClient)
	store := database.NewDBStore(context.Background(), mockClient)
	m := &users.User{}

	// Expect Get to be called
	mockClient.On("Get", mock.Anything, mock.MatchedBy(func(q sq.Sqlizer) bool {
		sql, args, _ := q.ToSql()
		return assert.Contains(t, sql, "WHERE email = ?") && assert.Equal(t, "test@test.com", args[0])
	}), mock.Anything).Return(nil)

	err := store.GetBy(m, "email", "test@test.com", m)
	assert.NoError(t, err)
}

func TestDBStore_Update(t *testing.T) {
	mockClient := new(database.MockDBClient)
	store := database.NewDBStore(context.Background(), mockClient)
	data := map[string]interface{}{"status": "archived"}

	m := &users.User{ID: 1}

	mockClient.On("Exec", mock.Anything, mock.MatchedBy(func(q sq.Sqlizer) bool {
		sql, _, _ := q.ToSql()
		return assert.Contains(t, sql, "UPDATE users") && assert.Contains(t, sql, "WHERE id = ?")
	})).Return(mockResult{}, nil)
	mockClient.On("Get", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := store.Update(m, data, m)
	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}