package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	HashedPassword string `json:"hashed_password"`
}

func NewPassword(userId, pwd string) (*Password, error) {
	hashedPwd, err := ToHash(pwd)
	if err != nil {
		return nil, err
	}
	return &Password{
		ID:             uuid.New().String(),
		UserID:         userId,
		HashedPassword: hashedPwd,
	}, nil
}

func ToHash(rawPwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (pwd *Password) Validate(rawPwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pwd.HashedPassword), []byte(rawPwd))
	if err != nil {
		return err
	}
	return nil
}
