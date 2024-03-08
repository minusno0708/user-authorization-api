package persistence

import (
	"database/sql"

	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type userPersistence struct {
	*sql.DB
}

func NewUserPersistence(db *sql.DB) repository.UserRepository {
	return &userPersistence{db}
}

func (up userPersistence) InsertUser(userID, username, password string) error {
	_, err := up.Exec(
		"INSERT INTO users (user_id, username, password) VALUES (?, ?, ?)",
		userID,
		username,
		password,
	)
	if err != nil {
		return err
	}

	return nil
}

func (up userPersistence) FindUserByUserID(userID string) (*domain.User, error) {
	user := domain.User{}
	err := up.QueryRow(
		"SELECT user_id, username, password FROM users WHERE user_id = ?",
		userID,
	).Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (up userPersistence) UpdateUsername(userID, username string) error {
	_, err := up.Exec(
		"UPDATE users SET username = ? WHERE user_id = ?",
		username,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (up userPersistence) DeleteUser(userID string) error {
	_, err := up.Exec(
		"DELETE FROM users WHERE user_id = ?",
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
