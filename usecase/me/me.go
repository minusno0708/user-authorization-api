package me

import (
	"user-register-api/domain/repository"
	"user-register-api/domain/token"
	"user-register-api/domain/user"
)

type MeUseCase interface {
	Get(tokenString string) (*user.User, error)
	Update(tokenString, username, email string) (*user.User, error)
	Delete(tokenString string) error
}

type meUseCase struct {
	secretKey string
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewMeUseCase(secretKey string, ur repository.UserRepository, tr repository.TokenRepository) MeUseCase {
	return &meUseCase{
		secretKey: secretKey,
		userRepo:  ur,
		tokenRepo: tr,
	}
}

func (mu meUseCase) Get(tokenString string) (*user.User, error) {
	err := mu.tokenRepo.IsValid(tokenString)
	if err != nil {
		return nil, err
	}

	token, err := token.ParseToken(tokenString, mu.secretKey)
	if err != nil {
		return nil, err
	}

	err = token.IsValid()
	if err != nil {
		return nil, err
	}

	user, err := mu.userRepo.FindByID(token.UserID())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mu meUseCase) Update(tokenString, username, email string) (*user.User, error) {
	err := mu.tokenRepo.IsValid(tokenString)
	if err != nil {
		return nil, err
	}

	token, err := token.ParseToken(tokenString, mu.secretKey)
	if err != nil {
		return nil, err
	}

	err = token.IsValid()
	if err != nil {
		return nil, err
	}

	updateUser, err := user.ReconstructUser(token.UserID(), username, email)
	if err != nil {
		return nil, err
	}

	err = mu.userRepo.Update(updateUser)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (mu meUseCase) Delete(tokenString string) error {
	err := mu.tokenRepo.IsValid(tokenString)
	if err != nil {
		return err
	}

	token, err := token.ParseToken(tokenString, mu.secretKey)
	if err != nil {
		return err
	}

	err = token.IsValid()
	if err != nil {
		return err
	}

	err = mu.userRepo.Delete(token.UserID())
	if err != nil {
		return err
	}

	return nil
}
