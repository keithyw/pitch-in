package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/keithyw/pitch-in/internal/config"
)

type DBClient interface {
	Get(ctx context.Context, builder squirrel.Sqlizer, dest interface{}) error
	Query(ctx context.Context, builder squirrel.Sqlizer, dest interface{}) error
	QueryMany(builder squirrel.Sqlizer) (*sql.Rows, error)
	Exec (ctx context.Context, builder squirrel.Sqlizer) (sql.Result, error)
}

type DBClientImpl struct {
	config *config.Config
	DB     *sqlx.DB
}

func NewDBClient(config *config.Config) (DBClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false&parseTime=true&interpolateParams=true",
		config.MysqlUser,
		config.MysqlPass,
		config.MysqlHost,
		config.MysqlDB,
	)
	fmt.Printf("dns: %s", dsn)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &DBClientImpl{
		config: config,
		DB: db,
	}, nil
}

func (s *DBClientImpl) Get(ctx context.Context, builder squirrel.Sqlizer, dest interface{}) error {
	qs, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	
	err = s.DB.GetContext(ctx, dest, qs, args...)
	if err != nil {
		return fmt.Errorf("failed to get: %s", err.Error())
	}
	return nil
}

func (s *DBClientImpl) Query(ctx context.Context, builder squirrel.Sqlizer, dest interface{}) error {
	qs, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	err = s.DB.SelectContext(ctx, dest, qs, args...)
	if err != nil {
		return fmt.Errorf("failed to select: %s", err.Error())
	}
	return nil
}

func (s *DBClientImpl) QueryMany(builder squirrel.Sqlizer) (*sql.Rows, error) {
	qs, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %s", err.Error())
	}
	return s.DB.Query(qs, args...)
}

func (s *DBClientImpl) Exec(ctx context.Context, builder squirrel.Sqlizer) (sql.Result, error) {
	qs, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %s", err.Error())
	}
	result, err := s.DB.ExecContext(ctx, qs, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to exec: %s", err.Error())
	}
	return result, nil
}
