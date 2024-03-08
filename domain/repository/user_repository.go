package repository

import (
	"user-register-api/domain"
)

type UserRepository interface {
	InsertUser(userID, username, password string) error
	FindUserByUserID(userID string) (*domain.User, error)
	UpdateUsername(userID string, username string) error
	DeleteUser(userID string) error
}
