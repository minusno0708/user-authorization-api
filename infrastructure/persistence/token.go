package persistence

import (
	"context"
	"time"

	"user-register-api/domain/repository"

	"github.com/redis/go-redis/v9"
)

type tokenPersistence struct {
	*redis.Client
}

func NewTokenPersistence(cdb *redis.Client) repository.TokenRepository {
	return &tokenPersistence{cdb}
}

func (tp tokenPersistence) SaveToken(userID, tokenUUID string) error {
	ctx := context.Background()

	err := tp.Set(ctx, userID, tokenUUID, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (tp tokenPersistence) ValidateToken(userID string) (string, error) {
	ctx := context.Background()

	tokenUUID, err := tp.Get(ctx, userID).Result()
	if err != nil {
		return "", err
	}

	err = tp.Expire(ctx, userID, time.Hour).Err()
	if err != nil {
		return "", err
	}

	return tokenUUID, nil
}

func (tp tokenPersistence) DeleteToken(userID string) error {
	ctx := context.Background()

	err := tp.Del(ctx, userID).Err()
	if err != nil {
		return err
	}

	return nil
}
