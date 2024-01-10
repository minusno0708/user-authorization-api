package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbAddress = "db:3306"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp("+dbAddress+")/test_db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}