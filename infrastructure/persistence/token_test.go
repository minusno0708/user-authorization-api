package persistence

import "testing"

func TestGenerateToken(t *testing.T) {
	tokenPersistence := NewTokenPersistence()

	exampleToken := "eyJhbG"
	err := tokenPersistence.GenerateToken(testUser.UserID, exampleToken)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateToken(t *testing.T) {
	tokenPersistence := NewTokenPersistence()
	exampleToken := "eyJhbG"

	token, err := tokenPersistence.ValidateToken(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if token != exampleToken {
		t.Errorf("Token is not match")
	}
}
