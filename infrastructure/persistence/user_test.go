package persistence

import (
	"testing"

	"user-register-api/config"
	"user-register-api/domain"
)

var testUser = domain.User{
	UserID:   "testuser_db",
	Username: "testuser_db",
	Password: "test_password",
}

func TestInsertUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence()

	err = userPersistence.InsertUser(db, testUser.UserID, testUser.Username, testUser.Password)
	if err != nil {
		t.Error(err)
	}
}

func TestInsertUserDuplicate(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence()

	err = userPersistence.InsertUser(db, testUser.UserID, testUser.Username, testUser.Password)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestFindUserByUserID(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence()

	user, err := userPersistence.FindUserByUserID(db, testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if user.UserID != testUser.UserID {
		t.Errorf("UserID is not match")
	}
	if user.Username != testUser.Username {
		t.Errorf("Username is not match")
	}
	if user.Password != testUser.Password {
		t.Errorf("Password is not match")
	}
}

func TestUpdateUsername(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence()

	updatedName := "testuser_db_updated"
	err = userPersistence.UpdateUsername(db, testUser.UserID, updatedName)
	if err != nil {
		t.Error(err)
	}

	user, err := userPersistence.FindUserByUserID(db, testUser.UserID)
	if err != nil {
		t.Error(err)
	}
	if user.Username != updatedName {
		t.Errorf("Failed to update username")
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence()

	err = userPersistence.DeleteUser(db, testUser.UserID)
	if err != nil {
		t.Error(err)
	}
}
