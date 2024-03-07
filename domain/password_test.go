package domain

import "testing"

func TestPassword(t *testing.T) {
	pwd := NewPassword("test_password1234")
	t.Log(pwd.value)
}
