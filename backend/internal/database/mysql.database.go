package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/keithyw/pitch-in/internal/config"
)

type MysqlDatabase struct {
	Config *config.Config
	DB *sql.DB
}

func NewMysqlDatabase(config *config.Config) (*MysqlDatabase, error) {
	mysqlConfig := mysql.Config{
		User: config.MysqlUser,
		Passwd: config.MysqlPass,
		Net: "tcp",
		Addr: config.MysqlHost,
		DBName: config.MysqlDB,
	}
	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &MysqlDatabase{
		Config: config,
		DB: db,
	}, nil
}