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

func (up userPersistence) InsertUser(user *domain.User) error {
	_, err := up.Exec(
		"INSERT INTO users (id, username, email) VALUES (?, ?, ?)",
		user.ID,
		user.Username,
		user.Email,
	)
	if err != nil {
		return err
	}

	return nil
}

func (up userPersistence) FindUserByID(userID string) (*domain.User, error) {
	user := domain.User{}
	err := up.QueryRow(
		"SELECT id, username, email FROM users WHERE id = ?",
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
	)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (up userPersistence) FindUserByUsername(username string) (*domain.User, error) {
	user := domain.User{}
	err := up.QueryRow(
		"SELECT id, username, email FROM users WHERE username = ?",
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
	)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (up userPersistence) UpdateUser(userID, username, email string) error {
	_, err := up.Exec(
		"UPDATE users SET username = ?, email = ? WHERE id = ?",
		username,
		email,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (up userPersistence) DeleteUser(userID string) error {
	_, err := up.Exec(
		"DELETE FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
