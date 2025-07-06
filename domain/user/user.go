package user

import (
	"user-register-api/pkg/errors"

	"user-register-api/pkg/uuid"
)

type User struct {
	id       string
	username string
	email    string
}

func NewUser(username, email string) (*User, error) {
	u := &User{}

	if err := u.SetID(uuid.New()); err != nil {
		return nil, err
	}
	if err := u.SetUsername(username); err != nil {
		return nil, err
	}
	if err := u.SetEmail(email); err != nil {
		return nil, err
	}
	return u, nil
}

func ReconstructUser(id, username, email string) (*User, error) {
	u := &User{}

	if err := u.SetID(id); err != nil {
		return nil, err
	}
	if err := u.SetUsername(username); err != nil {
		return nil, err
	}
	if err := u.SetEmail(email); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetID(id string) error {
	if !uuid.IsUUID(id) {
		return errors.ErrInternalServer
	}
	u.id = id
	return nil
}

func (u *User) SetUsername(username string) error {
	if username == "" {
		return errors.ErrBadRequest
	}
	u.username = username
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.ErrBadRequest
	}
	u.email = email
	return nil
}
