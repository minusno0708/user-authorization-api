package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Token struct {
	value *jwt.Token
}

var secretKey = "secret"

func NewToken(userID string) *Token {
	claims := jwt.MapClaims{
		"user_id": userID,
		"uuid":    uuid.New().String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return &Token{value: token}
}

func (t *Token) ToString() (string, error) {
	tokenString, err := t.value.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return &Token{value: token}, nil
}

func (t *Token) UUID() string {
	return t.value.Claims.(jwt.MapClaims)["uuid"].(string)
}

func (t *Token) UserID() string {
	return t.value.Claims.(jwt.MapClaims)["user_id"].(string)
}
