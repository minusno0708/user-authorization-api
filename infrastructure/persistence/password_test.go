package persistence

import (
	"testing"

	"user-register-api/config"
	"user-register-api/domain"
)

var rawPassword = "test_password"

var passwordTestUser = domain.NewUser(
	"testuser_password",
	"testuser_password",
)

func TestCreateUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence(db)

	err = userPersistence.InsertUser(passwordTestUser)
	if err != nil {
		t.Error(err)
	}

	user, err := userPersistence.FindUserByUsername(passwordTestUser.Username)
	if err != nil {
		t.Error(err)
	}
	passwordTestUser.ID = user.ID
}

func TestInsertPassword(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	passwordPersistence := NewPasswordPersistence(db)

	password, err := domain.NewPassword(
		passwordTestUser.ID,
		rawPassword,
	)
	if err != nil {
		t.Error(err)
	}

	err = passwordPersistence.InsertPassword(password)
	if err != nil {
		t.Error(err)
	}
}

func TestFindPasswordByUserID(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	passwordPersistence := NewPasswordPersistence(db)

	password, err := passwordPersistence.FindPasswordByUserID(passwordTestUser.ID)
	if err != nil {
		t.Error(err)
	}
	if password.UserID != passwordTestUser.ID {
		t.Errorf("UserID is not match")
	}
	if password.Validate(rawPassword) != nil {
		t.Errorf("Password is not match")
	}
}

func TestUpdatePassword(t *testing.T) {
	updatedPassword := "new_password"

	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	passwordPersistence := NewPasswordPersistence(db)

	updatedHashedPassword, err := domain.ToHash(updatedPassword)
	if err != nil {
		t.Error(err)
	}

	err = passwordPersistence.UpdatePassword(passwordTestUser.ID, updatedHashedPassword)
	if err != nil {
		t.Error(err)
	}

	password, err := passwordPersistence.FindPasswordByUserID(passwordTestUser.ID)
	if err != nil {
		t.Error(err)
	}

	if password.Validate(updatedPassword) != nil {
		t.Errorf("Password is not match")
	}
}

func TestDeletePassword(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := NewUserPersistence(db)
	passwordPersistence := NewPasswordPersistence(db)

	err = userPersistence.DeleteUser(passwordTestUser.ID)
	if err != nil {
		t.Error(err)
	}

	_, err = passwordPersistence.FindPasswordByUserID(passwordTestUser.ID)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
