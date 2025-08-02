package repository

import "user-register-api/domain/password"

type PasswordRepository interface {
	Insert(password *password.Password) error
	FindByUserID(userID string) (*password.Password, error)
	Update(password *password.Password) error
}
