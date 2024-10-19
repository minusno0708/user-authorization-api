package usecase

import (
	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type UserUseCase interface {
	InsertUser(username, email, password string) error
	FindUserByUserID(userID string) (*domain.User, error)
	UpdateUsername(userID, username, email string) error
	DeleteUser(userID string) error
	ValidateUser(username, password string) (string, error)
}

type userUseCase struct {
	userRepository     repository.UserRepository
	passwordRepository repository.PasswordRepository
}

func NewUserUseCase(ur repository.UserRepository, pr repository.PasswordRepository) UserUseCase {
	return &userUseCase{
		userRepository:     ur,
		passwordRepository: pr,
	}
}

func (uu userUseCase) InsertUser(username, email, rawPassword string) error {
	user := domain.NewUser(username, email)
	err := uu.userRepository.InsertUser(user)
	if err != nil {
		return err
	}

	password, err := domain.NewPassword(user.ID, rawPassword)
	if err != nil {
		return err
	}

	err = uu.passwordRepository.InsertPassword(password)
	if err != nil {
		return err
	}

	return nil
}

func (uu userUseCase) FindUserByUserID(userID string) (*domain.User, error) {
	user, err := uu.userRepository.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) FindUserByUsername(username string) (*domain.User, error) {
	user, err := uu.userRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) UpdateUsername(userID, updateUsername, updateEmail string) error {
	user, err := uu.userRepository.FindUserByID(userID)
	if err != nil {
		return err
	}

	if updateUsername == "" {
		updateUsername = user.Username
	}
	if updateEmail == "" {
		updateEmail = user.Email
	}

	err = uu.userRepository.UpdateUser(userID, updateUsername, updateEmail)
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

func (uu userUseCase) ValidateUser(username, password string) (string, error) {
	user, err := uu.userRepository.FindUserByUsername(username)
	if err != nil {
		return "", err
	}

	passwordData, err := uu.passwordRepository.FindPasswordByUserID(user.ID)
	if err != nil {
		return "", err
	}

	err = passwordData.Validate(password)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
