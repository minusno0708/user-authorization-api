package usecase

import (
	"database/sql"

	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type UserUseCase interface {
	InsertUser(db *sql.DB, userID, username, password string) (*domain.User, error)
	FindUserByUserID(db *sql.DB, userID string) (*domain.User, error)
	UpdateUsername(db *sql.DB, userID, username string) error
	DeleteUser(db *sql.DB, userID string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) InsertUser(db *sql.DB, userID, username, password string) (*domain.User, error) {
	if username == "" {
		username = userID
	}
	
	err := uu.userRepository.InsertUser(db, userID, username, password)
	if err != nil {
		return nil, err
	}

	user, err := uu.userRepository.FindUserByUserID(db, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu userUseCase) FindUserByUserID(db *sql.DB, userID string) (*domain.User, error) {
	user, err := uu.userRepository.FindUserByUserID(db, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) UpdateUsername(db *sql.DB, userID, username string) error {
	err := uu.userRepository.UpdateUsername(db, userID, username)
	if err != nil {
		return err
	}
	return nil
}

func (uu userUseCase) DeleteUser(db *sql.DB, userID string) error {
	err := uu.userRepository.DeleteUser(db, userID)
	if err != nil {
		return err
	}
	return nil
}