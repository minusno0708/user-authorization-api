package domain

import "testing"

var testUserID = "test_user_id"
var pwdString = "test_password1234"

func TestPasswordToHash(t *testing.T) {
	pwd, err := NewPassword(testUserID, pwdString)
	if err != nil {
		t.Error("Error while creating password")
	}
	if pwd.HashedPassword == "" {
		t.Error("Hashed password is empty")
	}
	if pwd.HashedPassword == pwdString {
		t.Error("password is not hashed")
	}
}

func TestPasswordIsMatch(t *testing.T) {
	pwd, err := NewPassword(testUserID, pwdString)
	if err != nil {
		t.Error("Error while creating password")
	}
	err = pwd.Validate(pwdString)
	if err != nil {
		t.Error("Password is not matched")
	}
}

func TestPasswordIsNotMatch(t *testing.T) {
	wrongPwd := "wrong_password"
	pwd, err := NewPassword(testUserID, pwdString)
	if err != nil {
		t.Error("Error while creating password")
	}
	err = pwd.Validate(wrongPwd)
	if err == nil {
		t.Error("Password is matched")
	}
}
