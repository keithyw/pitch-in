package config

import (
	"os"
	"strconv"
)

type Config struct {
	HttpPort string
	MysqlUser string
	MysqlPass string
	MysqlHost string
	MysqlDB string
	JWTExpirationTime int
	JWTSecretKey string
}

func NewConfig() *Config {
	exp, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		exp = 24
	}
	return &Config{
		HttpPort: os.Getenv("HTTP_PORT"),
		MysqlUser: os.Getenv("MYSQL_USER"),
		MysqlPass: os.Getenv("MYSQL_PASS"),
		MysqlHost: os.Getenv("MYSQL_HOST"),
		MysqlDB: os.Getenv("MYSQL_DATABASE"),
		JWTExpirationTime: exp,
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}