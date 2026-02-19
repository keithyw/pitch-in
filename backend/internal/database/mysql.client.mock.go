package database

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/mock"
)

type MockDBClient struct {
	mock.Mock
}

func (m *MockDBClient) Get(ctx context.Context, builder squirrel.Sqlizer, dest interface{}) error {
	args := m.Called(ctx, builder, dest)
	return args.Error(0)
}

func (m *MockDBClient) Query(ctx context.Context, builder squirrel.Sqlizer, dest interface{}) error {
	args := m.Called(ctx, builder, dest)
	return args.Error(0)
}

func (m *MockDBClient) QueryMany(builder squirrel.Sqlizer) (*sql.Rows, error) {
	args := m.Called(builder)
	return args.Get(0).(*sql.Rows), args.Error(1)
}

func (m *MockDBClient) Exec(ctx context.Context, builder squirrel.Sqlizer) (sql.Result, error) {
	args := m.Called(ctx, builder)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(sql.Result), args.Error(1)
}