package persistence

import (
	"testing"
	"user-register-api/config"
)

var exampleUUID = "de5fe5d7-eec2-4fba-e071-fa2de7c1e440"

func TestSaveToken(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := NewTokenPersistence(cdb)

	err = tokenPersistence.SaveToken(testUser.UserID, exampleUUID)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateToken(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := NewTokenPersistence(cdb)

	token, err := tokenPersistence.ValidateToken(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if token != exampleUUID {
		t.Errorf("UUID is not match")
	}
}

func TestDeleteToken(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := NewTokenPersistence(cdb)

	err = tokenPersistence.DeleteToken(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
}

func TestValidateTokenDeleted(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := NewTokenPersistence(cdb)

	_, err = tokenPersistence.ValidateToken(testUser.UserID)
	if err == nil {
		t.Error("Token is not deleted")
	}
}
