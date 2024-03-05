package usecase

import (
	"errors"
	"time"
	"user-register-api/domain/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenUseCase interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (string, error)
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

func (au tokenUseCase) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	userID := token.Claims.(jwt.MapClaims)["user_id"].(string)

	return userID, nil
}
