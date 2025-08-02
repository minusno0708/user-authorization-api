package config

import (
	"testing"
)

func TestConnectionDB(t *testing.T) {
	db, err := ConnectDB("localhost:3306")
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()
}

func TestConnectionCacheDB(t *testing.T) {
	_, err := ConnectCacheDB("localhost:6379")
	if err != nil {
		t.Fatalf("Error connecting to the database: %v", err)
	}
}
