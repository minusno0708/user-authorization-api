package config

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
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

func ConnectCacheDB() (*redis.Client, error) {
	var ctx = context.Background()

	cdb := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "",
		DB:       0,
	})
	_, err := cdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return cdb, nil
}
