package usecase

import (
	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type UserUseCase interface {
	InsertUser(userID, username, password string) (*domain.User, error)
	FindUserByUserID(userID string) (*domain.User, error)
	UpdateUsername(userID, username string) (*domain.User, error)
	DeleteUser(userID string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) InsertUser(userID, username, password string) (*domain.User, error) {
	if username == "" {
		username = userID
	}

	hashPassword, err := domain.NewPassword(password).ToHash()
	if err != nil {
		return nil, err
	}

	err = uu.userRepository.InsertUser(userID, username, hashPassword)
	if err != nil {
		return nil, err
	}

	user, err := uu.userRepository.FindUserByUserID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu userUseCase) FindUserByUserID(userID string) (*domain.User, error) {
	user, err := uu.userRepository.FindUserByUserID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) UpdateUsername(userID, username string) (*domain.User, error) {
	err := uu.userRepository.UpdateUsername(userID, username)
	if err != nil {
		return nil, err
	}

	user, err := uu.userRepository.FindUserByUserID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu userUseCase) DeleteUser(userID string) error {
	err := uu.userRepository.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}
