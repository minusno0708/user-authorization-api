package usecase

import (
	"testing"
	"user-register-api/config"
	"user-register-api/infrastructure/persistence"
)

var passwordTestUserID = "00000000-0000-0000-0000-000000000000"

func TestUsecaseGenerateToken(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := persistence.NewTokenPersistence(cdb)
	tokenUseCase := NewTokenUseCase(tokenPersistence)

	_, err = tokenUseCase.GenerateToken(passwordTestUserID)
	if err != nil {
		t.Error(err)
	}
}

func TestUsecaseValidateToken(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := persistence.NewTokenPersistence(cdb)
	tokenUseCase := NewTokenUseCase(tokenPersistence)

	token, err := tokenUseCase.GenerateToken(passwordTestUserID)
	if err != nil {
		t.Error(err)
	}

	userID, err := tokenUseCase.ValidateToken(token)
	if err != nil {
		t.Error(err)
	}
	if userID != passwordTestUserID {
		t.Errorf("UserID is not correct")
	}
}

func TestUsecaseValidateInvalidToken(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := persistence.NewTokenPersistence(cdb)
	tokenUseCase := NewTokenUseCase(tokenPersistence)

	token := "invalid_token"

	_, err = tokenUseCase.ValidateToken(token)
	if err == nil {
		t.Errorf("Invalid token is accepted")
	}
}

func TestUsecaseDeleteToken(t *testing.T) {
	cdb, err := config.ConnectCacheDB()
	if err != nil {
		t.Error(err)
	}

	tokenPersistence := persistence.NewTokenPersistence(cdb)
	tokenUseCase := NewTokenUseCase(tokenPersistence)

	token, err := tokenUseCase.GenerateToken(passwordTestUserID)
	if err != nil {
		t.Error(err)
	}

	err = tokenUseCase.DeleteToken(token)
	if err != nil {
		t.Error(err)
	}

	_, err = tokenUseCase.ValidateToken(token)
	if err == nil {
		t.Errorf("Token is not deleted")
	}
}
