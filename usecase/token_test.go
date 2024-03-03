package usecase

import (
	"testing"
	"user-register-api/infrastructure/persistence"
)

func TestUsecaseGenerateToken(t *testing.T) {
	tokenPersistence := persistence.NewTokenPersistence()
	tokenUseCase := NewTokenUseCase(tokenPersistence)

	_, err := tokenUseCase.GenerateToken(testUser.UserID)
	if err != nil {
		t.Error(err)
	}
}

func TestUsecaseValidateToken(t *testing.T) {
	tokenPersistence := persistence.NewTokenPersistence()
	tokenUseCase := NewTokenUseCase(tokenPersistence)

	token, err := tokenUseCase.GenerateToken(testUser.UserID)
	if err != nil {
		t.Error(err)
	}

	userID, err := tokenUseCase.ValidateToken(token)
	if err != nil {
		t.Error(err)
	}
	if userID != testUser.UserID {
		t.Errorf("UserID is not correct")
	}
}

func TestUsecaseValidateInvalidToken(t *testing.T) {
	tokenPersistence := persistence.NewTokenPersistence()
	tokenUseCase := NewTokenUseCase(tokenPersistence)

	token := "invalid_token"

	_, err := tokenUseCase.ValidateToken(token)
	if err == nil {
		t.Errorf("Invalid token is accepted")
	}
}
