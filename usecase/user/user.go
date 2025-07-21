package user

import (
	"user-register-api/domain/password"
	"user-register-api/domain/repository"
	"user-register-api/domain/user"
)

type UserUseCase interface {
	Create(username, email, password string) (*user.User, error)
	FindByUsername(username string) (*user.User, error)
	Update(userID, username, email string) (*user.User, error)
	Delete(userID string) error
	Validate(username, password string) (string, error)
}

type userUseCase struct {
	userRepo     repository.UserRepository
	passwordRepo repository.PasswordRepository
}

func NewUserUseCase(ur repository.UserRepository, pr repository.PasswordRepository) UserUseCase {
	return &userUseCase{
		userRepo:     ur,
		passwordRepo: pr,
	}
}

func (uu userUseCase) Create(username, email, rawPassword string) (*user.User, error) {
	user, err := user.NewUser(username, email)
	if err != nil {
		return nil, err
	}

	userPassword, err := password.NewPassword(user.ID(), rawPassword)
	if err != nil {
		return nil, err
	}

	err = uu.userRepo.Insert(user)
	if err != nil {
		return nil, err
	}

	err = uu.passwordRepo.Insert(userPassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu userUseCase) FindByUsername(username string) (*user.User, error) {
	user, err := uu.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) Update(userID, updateUsername, updateEmail string) (*user.User, error) {
	user, err := user.ReconstructUser(userID, updateUsername, updateEmail)
	if err != nil {
		return nil, err
	}

	err = uu.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu userUseCase) Delete(userID string) error {
	err := uu.userRepo.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}

func (uu userUseCase) Validate(username, password string) (string, error) {
	user, err := uu.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	userPassword, err := uu.passwordRepo.FindByUserID(user.ID())
	if err != nil {
		return "", err
	}

	err = userPassword.Compare(password)
	if err != nil {
		return "", err
	}

	return user.ID(), nil
}
