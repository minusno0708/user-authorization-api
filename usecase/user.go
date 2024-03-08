package usecase

import (
	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type UserUseCase interface {
	InsertUser(userID, username, password string) error
	FindUserByUserID(userID string) (*domain.User, error)
	UpdateUsername(userID, username string) error
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

func (uu userUseCase) InsertUser(userID, username, password string) error {
	hashPassword, err := domain.NewPassword(password).ToHash()
	if err != nil {
		return err
	}

	user := domain.NewUser(userID, username, hashPassword)

	err = uu.userRepository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu userUseCase) FindUserByUserID(userID string) (*domain.User, error) {
	user, err := uu.userRepository.FindUserByUserID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) UpdateUsername(userID, username string) error {
	err := uu.userRepository.UpdateUsername(userID, username)
	if err != nil {
		return err
	}

	return nil
}

func (uu userUseCase) DeleteUser(userID string) error {
	err := uu.userRepository.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}
