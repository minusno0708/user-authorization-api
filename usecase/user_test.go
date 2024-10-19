package usecase

import (
	"testing"

	"user-register-api/config"
	"user-register-api/infrastructure/persistence"
)

type RequestUser struct {
	Username string
	Email    string
	Password string
}

var testUser = RequestUser{
	Username: "testuser_usecase",
	Email:    "testuser_email",
	Password: "testuser_password",
}

var testUserID string

func TestUsecaseInsertUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	userPersistence := persistence.NewUserPersistence(db)
	passwordPersintence := persistence.NewPasswordPersistence(db)
	userUseCase := NewUserUseCase(userPersistence, passwordPersintence)

	err = userUseCase.InsertUser(testUser.Username, testUser.Email, testUser.Password)
	if err != nil {
		t.Error(err)
	}

	testUserID, err = userUseCase.ValidateUser(testUser.Username, testUser.Password)
	if err != nil {
		t.Error(err)
	}

	user, err := userUseCase.FindUserByUserID(testUserID)
	if err != nil {
		t.Error(err)
	}

	if user.Username != testUser.Username {
		t.Errorf("Username is not correct")
	}
	if user.Email != testUser.Email {
		t.Errorf("Email is not correct")
	}
}

func TestUsecaseUpdateUser(t *testing.T) {
	updateUsername := "updated_username"
	updateEmail := "updated_email"

	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	userPersistence := persistence.NewUserPersistence(db)
	passwordPersintence := persistence.NewPasswordPersistence(db)
	userUseCase := NewUserUseCase(userPersistence, passwordPersintence)

	user, err := userUseCase.FindUserByUserID(testUserID)
	if err != nil {
		t.Error(err)
	}

	err = userUseCase.UpdateUsername(user.ID, updateUsername, updateEmail)
	if err != nil {
		t.Error(err)
	}

	updatedUser, err := userUseCase.FindUserByUserID(user.ID)
	if err != nil {
		t.Error(err)
	}

	if updatedUser.Username != updateUsername {
		t.Errorf("Username is not correct")
	}
	if updatedUser.Email != updateEmail {
		t.Errorf("Email is not correct")
	}
}

func TestUsecaseDeleteUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	userPersistence := persistence.NewUserPersistence(db)
	passwordPersintence := persistence.NewPasswordPersistence(db)
	userUseCase := NewUserUseCase(userPersistence, passwordPersintence)

	err = userUseCase.DeleteUser(testUserID)
	if err != nil {
		t.Error(err)
	}

	_, err = userUseCase.FindUserByUserID(testUserID)
	if err == nil {
		t.Errorf("User is not deleted")
	}
}
