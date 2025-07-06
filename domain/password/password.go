package password

import (
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	id             string
	userID         string
	hashedPassword string
}

func NewPassword(userID, password string) (*Password, error) {
	p := &Password{}

	if err := p.SetID(uuid.New()); err != nil {
		return nil, err
	}
	if err := p.SetUserID(userID); err != nil {
		return nil, err
	}
	if err := p.SetPassword(password); err != nil {
		return nil, err
	}

	return p, nil
}

func ReconstructPassword(id, userID, hashedPassword string) (*Password, error) {
	p := &Password{}

	if err := p.SetID(id); err != nil {
		return nil, err
	}
	if err := p.SetUserID(userID); err != nil {
		return nil, err
	}
	if err := p.SetHashedPassword(hashedPassword); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Password) ID() string {
	return p.id
}

func (p *Password) UserID() string {
	return p.userID
}

func (p *Password) HashedPassword() string {
	return p.hashedPassword
}

func (p *Password) SetID(id string) error {
	if !uuid.IsUUID(id) {
		return errors.ErrInternalServer
	}

	p.id = id

	return nil
}

func (p *Password) SetUserID(userID string) error {
	if !uuid.IsUUID(userID) {
		return errors.ErrInternalServer
	}

	p.userID = userID

	return nil
}

func (p *Password) SetPassword(password string) error {
	if password == "" {
		return errors.ErrBadRequest
	}
	if len(password) < 8 {
		return errors.ErrBadRequest
	}

	hashedPassword, err := hash(password)
	if err != nil {
		return errors.ErrInternalServer
	}
	p.hashedPassword = hashedPassword

	return nil
}

func (p *Password) SetHashedPassword(hashedPassword string) error {
	if hashedPassword == "" {
		return errors.ErrBadRequest
	}

	p.hashedPassword = hashedPassword

	return nil
}

func (p *Password) Compare(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.hashedPassword), []byte(password))
	if err != nil {
		return errors.ErrUnauthorized
	}
	return nil
}

func hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
