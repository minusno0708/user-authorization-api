package token

import (
	"time"

	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"

	"github.com/golang-jwt/jwt"
)

const (
	lifetime = time.Hour * 24
)

type Token struct {
	value *jwt.Token
}

func NewToken(userID string) *Token {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"uuid":    uuid.New(),
			"exp":     time.Now().Add(lifetime).Unix(),
		},
	)

	return &Token{value: token}
}

func (t *Token) SignedString(secretKey string) (string, error) {
	tokenString, err := t.value.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.ErrInternalServer
	}
	return tokenString, nil
}

func ParseToken(tokenString, secretKey string) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrUnauthorized
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.ErrUnauthorized
	}
	return &Token{value: token}, nil
}

func (t *Token) UUID() string {
	return t.value.Claims.(jwt.MapClaims)["uuid"].(string)
}

func (t *Token) UserID() string {
	return t.value.Claims.(jwt.MapClaims)["user_id"].(string)
}

func (t *Token) Exp() int64 {
	exp, ok := t.value.Claims.(jwt.MapClaims)["exp"].(int64)
	if !ok {
		exp = int64(t.value.Claims.(jwt.MapClaims)["exp"].(float64))
	}
	return exp
}

func (t *Token) IsValid() error {
	if time.Now().Unix() > int64(t.Exp()) {
		return errors.ErrUnauthorized
	}

	return nil
}
