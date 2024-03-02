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

	err = cdb.Set(ctx, token, userID, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (tp tokenPersistence) ValidateToken(token string) (string, error) {
	ctx := context.Background()

	cdb, err := config.ConnectCacheDB()
	if err != nil {
		return "", err
	}

	userID, err := cdb.Get(ctx, token).Result()
	if err != nil {
		return "", err
	}

	return userID, nil
}
