package usecase

import (
	"testing"

	"user-register-api/config"
	"user-register-api/domain"
	"user-register-api/infrastructure/persistence"
)

var testUser = domain.User{
	UserID:   "testuser_usecase",
	Password: "test_password",
}

func TestUsecaseInsertUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := persistence.NewUserPersistence(db)
	userUseCase := NewUserUseCase(userPersistence)

	err = userUseCase.InsertUser(testUser.UserID, testUser.Username, testUser.Password)
	if err != nil {
		t.Error(err)
	}

	user, err := userUseCase.FindUserByUserID(testUser.UserID)
	if err != nil {
		t.Error(err)
	}

	if user.Username != testUser.UserID {
		t.Errorf("Username is not correct")
	}
}

func TestUsecaseUpdateUser(t *testing.T) {
	updatedUsername := "updated_username"

	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := persistence.NewUserPersistence(db)
	userUseCase := NewUserUseCase(userPersistence)

	err = userUseCase.UpdateUsername(testUser.UserID, updatedUsername)
	if err != nil {
		t.Error(err)
	}

	user, err := userUseCase.FindUserByUserID(testUser.UserID)
	if err != nil {
		t.Error(err)
	}

	if user.Username != updatedUsername {
		t.Errorf("Username is not correct")
	}
}

func TestUsecaseDeleteUser(t *testing.T) {
	db, err := config.ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	userPersistence := persistence.NewUserPersistence(db)
	userUseCase := NewUserUseCase(userPersistence)

	err = userUseCase.DeleteUser(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
}
