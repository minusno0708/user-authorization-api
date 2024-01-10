package repository

import (
	"database/sql"

	"user-register-api/domain"
)

type UserRepository interface {
	InsertUser(db *sql.DB, userID, username, password string) error
	FindUserByUserID(db *sql.DB, userID string) (*domain.User, error)
	UpdateUsername(db *sql.DB, userID string, username string) error
	DeleteUser(db *sql.DB, userID string) error
}