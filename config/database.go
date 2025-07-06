package config

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp("+dsn+")/test_db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectCacheDB(dsn string) (*redis.Client, error) {
	var ctx = context.Background()

	cdb := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})
	_, err := cdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return cdb, nil
}
