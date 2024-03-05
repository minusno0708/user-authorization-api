package usecase

import (
	"time"
	"user-register-api/domain/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

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
	claims := jwt.MapClaims{
		"user_id": userID,
		"uuid":    uuid.New().String(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	err = au.tokenRepository.GenerateToken(userID, claims["uuid"].(string))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (au tokenUseCase) ValidateToken(token string) (string, error) {
	userID, err := au.tokenRepository.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}
