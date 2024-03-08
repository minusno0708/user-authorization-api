package persistence

import (
	"context"
	"time"
	"user-register-api/config"

	"user-register-api/domain/repository"
)

type tokenPersistence struct{}

func NewTokenPersistence() repository.TokenRepository {
	return &tokenPersistence{}
}

func (tp tokenPersistence) SaveToken(userID, tokenUUID string) error {
	ctx := context.Background()

	cdb, err := config.ConnectCacheDB()
	if err != nil {
		return err
	}

	err = cdb.Set(ctx, userID, tokenUUID, time.Hour).Err()
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

	tokenUUID, err := cdb.Get(ctx, userID).Result()
	if err != nil {
		return "", err
	}

	err = cdb.Expire(ctx, userID, time.Hour).Err()
	if err != nil {
		return "", err
	}

	return tokenUUID, nil
}

func (tp tokenPersistence) DeleteToken(userID string) error {
	ctx := context.Background()

	cdb, err := config.ConnectCacheDB()
	if err != nil {
		return err
	}

	err = cdb.Del(ctx, userID).Err()
	if err != nil {
		return err
	}

	return nil
}
