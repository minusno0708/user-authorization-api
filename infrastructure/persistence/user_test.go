package persistence

import (
	"testing"

	"user-register-api/config"
	"user-register-api/domain"
)

var testUser = domain.User{
	Username: "testuser_name",
	Email:    "testuser_email",
}

func TestInsertUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence(db)

	err = userPersistence.InsertUser(&testUser)
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
	userPersistence := NewUserPersistence(db)

	err = userPersistence.InsertUser(&testUser)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestFindUserByUsername(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence(db)

	user, err := userPersistence.FindUserByUsername(testUser.Username)
	if err != nil {
		t.Error(err)
	}
	if user.Username != testUser.Username {
		t.Errorf("Username is not match")
	}
	if user.Email != testUser.Email {
		t.Errorf("Password is not match")
	}
	testUser.ID = user.ID
}

func TestFindUserByID(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence(db)

	user, err := userPersistence.FindUserByID(testUser.ID)
	if err != nil {
		t.Error(err)
	}
	if user.Username != testUser.Username {
		t.Errorf("Username is not match")
	}
	if user.Email != testUser.Email {
		t.Errorf("Password is not match")
	}
}

func TestUpdateUsername(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence(db)

	updatedName := "testuser_name_updated"
	updatedEmail := "testuser_email_updated"
	err = userPersistence.UpdateUser(testUser.ID, updatedName, updatedEmail)
	if err != nil {
		t.Error(err)
	}

	user, err := userPersistence.FindUserByID(testUser.ID)
	if err != nil {
		t.Error(err)
	}
	if user.ID != testUser.ID {
		t.Errorf("Unmatched user ID")
	}
	if user.Username != updatedName {
		t.Errorf("Failed to update username")
	}
	if user.Email != updatedEmail {
		t.Errorf("Failed to update email")
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence(db)

	err = userPersistence.DeleteUser(testUser.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = userPersistence.FindUserByUsername(testUser.Username)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
