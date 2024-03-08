package usecase

import (
	"errors"
	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type TokenUseCase interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (string, error)
	DeleteToken(tokenString string) error
}

type tokenUseCase struct {
	tokenRepository repository.TokenRepository
}

func NewTokenUseCase(tr repository.TokenRepository) TokenUseCase {
	return &tokenUseCase{
		tokenRepository: tr,
	}
}

func (tu tokenUseCase) GenerateToken(userID string) (string, error) {
	token := domain.NewToken(userID)

	tokenString, err := token.ToString()
	if err != nil {
		return "", err
	}

	err = tu.tokenRepository.SaveToken(userID, token.UUID())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tu tokenUseCase) ValidateToken(tokenString string) (string, error) {
	token, err := domain.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	if token.IsExpired() {
		return "", errors.New("token is expired")
	}

	userID := token.UserID()
	tokenUUID, err := tu.tokenRepository.ValidateToken(userID)
	if err != nil {
		return "", err
	}

	if tokenUUID != token.UUID() {
		return "", errors.New("invalid token")
	}

	return userID, nil
}

func (tu tokenUseCase) DeleteToken(tokenString string) error {
	token, err := domain.ParseToken(tokenString)
	if err != nil {
		return err
	}

	userID := token.UserID()
	err = tu.tokenRepository.DeleteToken(userID)
	if err != nil {
		return err
	}

	return err
}
