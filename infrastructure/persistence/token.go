package persistence

import (
	"context"
	"user-register-api/config"

	"user-register-api/domain/repository"
)

type tokenPersistence struct{}

func NewTokenPersistence() repository.TokenRepository {
	return &tokenPersistence{}
}

func (tp tokenPersistence) GenerateToken(userID, token string) error {
	ctx := context.Background()

	cdb, err := config.ConnectCacheDB()
	if err != nil {
		return err
	}

	err = cdb.Set(ctx, userID, token, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (tp tokenPersistence) ValidateToken(userID string) (string, error) {
	ctx := context.Background()

	cdb, err := config.ConnectCacheDB()
	if err != nil {
		return "", err
	}

	token, err := cdb.Get(ctx, userID).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}
