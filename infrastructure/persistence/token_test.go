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
	userID, err := tokenPersistence.ValidateToken(exampleToken)
	if err != nil {
		t.Error(err)
	}
	if userID != testUser.UserID {
		t.Errorf("UserID is not match")
	}
}
