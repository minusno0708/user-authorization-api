package domain

import "testing"

var pwdString = "test_password1234"

var hashPwd string

func TestPasswordToHash(t *testing.T) {
	pwd := NewPassword(pwdString)
	hash, err := pwd.ToHash()
	if err != nil {
		t.Error("Error while hashing password")
	}
	if hash == "" {
		t.Error("Hashing password is empty")
	}
	if hash == pwd.value {
		t.Error("Hashing password is not hashed")
	}
	hashPwd = hash
}

func TestPasswordIsMatch(t *testing.T) {
	pwd := NewPassword(pwdString)
	err := pwd.Compare(hashPwd)
	if err != nil {
		t.Error("Password is not matched")
	}
}

func TestPasswordIsNotMatch(t *testing.T) {
	pwd := NewPassword("wrong_password")
	err := pwd.Compare(hashPwd)
	if err == nil {
		t.Error("Password is matched")
	}
}
