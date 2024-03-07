package persistence

import "testing"

var exampleUuid = "de5fe5d7-eec2-4fba-e071-fa2de7c1e440"

func TestSaveToken(t *testing.T) {
	tokenPersistence := NewTokenPersistence()

	err := tokenPersistence.SaveToken(testUser.UserID, exampleUuid)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateToken(t *testing.T) {
	tokenPersistence := NewTokenPersistence()

	token, err := tokenPersistence.ValidateToken(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if token != exampleUuid {
		t.Errorf("UUID is not match")
	}
}

func TestDeleteToken(t *testing.T) {
	tokenPersistence := NewTokenPersistence()

	err := tokenPersistence.DeleteToken(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateTokenDeleted(t *testing.T) {
	tokenPersistence := NewTokenPersistence()

	_, err := tokenPersistence.ValidateToken(testUser.UserID)
	if err == nil {
		t.Error("Token is not deleted")
	}
}
