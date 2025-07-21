package token

import (
	"context"
	"time"

	"user-register-api/domain/repository"
	"user-register-api/pkg/errors"

	"github.com/redis/go-redis/v9"
)

type tokenRepository struct {
	*redis.Client
}

func NewTokenRepository(cdb *redis.Client) repository.TokenRepository {
	return &tokenRepository{cdb}
}

func (tr tokenRepository) Invalidate(tokenString string) error {
	ctx := context.Background()

	err := tr.Set(ctx, tokenString, "invalid", time.Hour*72).Err()
	if err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

func (tr tokenRepository) IsValid(tokenString string) error {
	ctx := context.Background()

	_, err := tr.Get(ctx, tokenString).Result()
	if err != nil {
		return nil
	}

	return errors.ErrForbidden
}
