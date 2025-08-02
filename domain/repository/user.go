package repository

import "user-register-api/domain/user"

type UserRepository interface {
	Insert(user *user.User) error
	FindByID(userID string) (*user.User, error)
	FindByUsername(username string) (*user.User, error)
	Update(user *user.User) error
	Delete(userID string) error
}
