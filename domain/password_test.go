package domain

import "testing"

func TestPasswordToHash(t *testing.T) {
	pwd := NewPassword("test_password1234")
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
}
