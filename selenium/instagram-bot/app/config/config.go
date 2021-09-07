package config

import (
	"os"
)

// Config returns a map with environment variables values.
func Config() map[string]string {

	setEnv()

	cfg := make(map[string]string)

	cfg["MYSQL_DB_ENDPOINT"] = os.Getenv("MYSQL_DB_ENDPOINT")
	cfg["MYSQL_DB_PORT"] = os.Getenv("MYSQL_DB_PORT")
	cfg["MYSQL_DB_SCHEMA"] = os.Getenv("MYSQL_DB_SCHEMA")
	cfg["MYSQL_DB_USER"] = os.Getenv("MYSQL_DB_USER")
	cfg["MYSQL_DB_PASSWORD"] = os.Getenv("MYSQL_DB_PASSWORD")

	return cfg
}

// TODO: Coletar esses dados de um arquivo ENV
func setEnv() {
	os.Setenv("MYSQL_DB_ENDPOINT", "db-mysql:3306")
	os.Setenv("MYSQL_DB_PORT", "3306")
	os.Setenv("MYSQL_DB_SCHEMA", "instagram_bot")
	os.Setenv("MYSQL_DB_USER", "root")
	os.Setenv("MYSQL_DB_PASSWORD", "root")
}
