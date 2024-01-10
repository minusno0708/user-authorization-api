package persistence

import (
	"database/sql"

	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) InsertUser(db *sql.DB, userID, username, password string) error {
	_, err := db.Exec(
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

func (up userPersistence) FindUserByUserID(db *sql.DB, userID string) (*domain.User, error) {
	user := domain.User{}
	err := db.QueryRow(
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

func (up userPersistence) UpdateUsername(db *sql.DB, userID, username string) error {
	_, err := db.Exec(
		"UPDATE users SET username = ? WHERE user_id = ?",
		username,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (up userPersistence) DeleteUser(db *sql.DB, userID string) error {
	_, err := db.Exec(
		"DELETE FROM users WHERE user_id = ?",
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}