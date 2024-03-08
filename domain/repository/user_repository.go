package repository

import (
	"user-register-api/domain"
)

type UserRepository interface {
	InsertUser(user *domain.User) error
	FindUserByUserID(userID string) (*domain.User, error)
	UpdateUsername(userID string, username string) error
	DeleteUser(userID string) error
}
