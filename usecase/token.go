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

var secretKey = "secret"

func tokenParse(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (tu tokenUseCase) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"uuid":    uuid.New().String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	err = tu.tokenRepository.SaveToken(userID, claims["uuid"].(string))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tu tokenUseCase) ValidateToken(tokenString string) (string, error) {
	token, err := tokenParse(tokenString)
	if err != nil {
		return "", err
	}

	userID := token.Claims.(jwt.MapClaims)["user_id"].(string)
	tokenUuid, err := tu.tokenRepository.ValidateToken(userID)
	if err != nil {
		return "", err
	}

	if tokenUuid != token.Claims.(jwt.MapClaims)["uuid"].(string) {
		return "", errors.New("invalid token")
	}

	return userID, nil
}

func (tu tokenUseCase) DeleteToken(tokenString string) error {
	token, err := tokenParse(tokenString)
	if err != nil {
		return err
	}

	userID := token.Claims.(jwt.MapClaims)["user_id"].(string)
	err = tu.tokenRepository.DeleteToken(userID)
	if err != nil {
		return err
	}

	return err
}
