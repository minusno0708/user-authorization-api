package usecase

import (
	"user-register-api/domain/repository"
	"user-register-api/domain/token"
	"user-register-api/pkg/errors"
)

type AuthUseCase interface {
	Login(username, rawPassword string) (string, error)
	Logout(tokenString string) error
}

type authUseCase struct {
	secretKey    string
	userRepo     repository.UserRepository
	passwordRepo repository.PasswordRepository
	tokenRepo    repository.TokenRepository
}

func NewAuthUseCase(secretKey string, ur repository.UserRepository, pr repository.PasswordRepository, tr repository.TokenRepository) AuthUseCase {
	return &authUseCase{
		secretKey:    secretKey,
		userRepo:     ur,
		passwordRepo: pr,
		tokenRepo:    tr,
	}
}

func (au authUseCase) Login(username, rawPassword string) (string, error) {
	user, err := au.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.ErrUnauthorized
	}

	password, err := au.passwordRepo.FindByUserID(user.ID())
	if err != nil {
		return "", err
	}

	err = password.Compare(rawPassword)
	if err != nil {
		return "", err
	}

	token := token.NewToken(user.ID())

	tokenString, err := token.SignedString(au.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (au authUseCase) Logout(tokenString string) error {
	token, err := token.ParseToken(tokenString, au.secretKey)
	if err != nil {
		return err
	}

	err = token.IsValid()
	if err != nil {
		return err
	}

	err = au.tokenRepo.Invalidate(tokenString)
	if err != nil {
		return err
	}

	return nil
}
