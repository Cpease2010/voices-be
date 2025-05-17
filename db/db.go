package db

import (
	"database/sql"
	"fmt"
	"voices/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	user := config.GetEnv("MYSQL_USER", "voices_user")
	pass := config.GetEnv("MYSQL_PASSWORD", "voices_pass")
	host := config.GetEnv("MYSQL_HOST", "localhost")
	port := config.GetEnv("MYSQL_PORT", "3306")
	name := config.GetEnv("MYSQL_DATABASE", "voices")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	var err error
	DB, err = sql.Open("mysql", dsn)
	return err
}
