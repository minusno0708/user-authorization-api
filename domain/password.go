package domain

import "golang.org/x/crypto/bcrypt"

type Password struct {
	value string
}

func NewPassword(pwd string) *Password {
	return &Password{value: pwd}
}

func (pwd *Password) ToHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd.value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (pwd *Password) IsMatch(hashPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd.value))
	if err != nil {
		return false, err
	}
	return true, nil
}
