package config

import "os"

type Config struct {
	HttpPort string
	MysqlUser string
	MysqlPass string
	MysqlHost string
	MysqlDB string
}

func NewConfig() *Config {
	return &Config{
		HttpPort: os.Getenv("HTTP_PORT"),
		MysqlUser: os.Getenv("MYSQL_USER"),
		MysqlPass: os.Getenv("MYSQL_PASS"),
		MysqlHost: os.Getenv("MYSQL_HOST"),
		MysqlDB: os.Getenv("MYSQL_DATABASE"),
	}
}