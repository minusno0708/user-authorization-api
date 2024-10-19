package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	IsDeleted bool   `json:"is_deleted"`
}

func NewUser(username, email string) *User {
	return &User{
		ID:       uuid.New().String(),
		Username: username,
		Email:    email,
	}
}
