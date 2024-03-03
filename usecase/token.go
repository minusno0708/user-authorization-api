package usecase

import "user-register-api/domain/repository"

type TokenUseCase interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (string, error)
}

type tokenUseCase struct {
	tokenRepository repository.TokenRepository
}

func NewTokenUseCase(tr repository.TokenRepository) TokenUseCase {
	return &tokenUseCase{
		tokenRepository: tr,
	}
}

func (au tokenUseCase) GenerateToken(userID string) (string, error) {
	token := "example_token"
	err := au.tokenRepository.GenerateToken(userID, token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (au tokenUseCase) ValidateToken(token string) (string, error) {
	userID, err := au.tokenRepository.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}
