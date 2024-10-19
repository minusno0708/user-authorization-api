package repository

import (
	"user-register-api/domain"
)

type PasswordRepository interface {
	InsertPassword(user *domain.Password) error
	FindPasswordByUserID(userID string) (*domain.Password, error)
	UpdatePassword(userID, updatePassword string) error
}
