package repository

import (
	"user-register-api/domain"
)

type UserRepository interface {
	InsertUser(user *domain.User) error
	FindUserByID(userID string) (*domain.User, error)
	FindUserByUsername(username string) (*domain.User, error)
	UpdateUser(usreID, username, email string) error
	DeleteUser(userID string) error
}
