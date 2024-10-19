package persistence

import (
	"database/sql"

	"user-register-api/domain"
	"user-register-api/domain/repository"
)

type passwordPersistence struct {
	*sql.DB
}

func NewPasswordPersistence(db *sql.DB) repository.PasswordRepository {
	return &passwordPersistence{db}
}

func (pp passwordPersistence) InsertPassword(password *domain.Password) error {
	_, err := pp.Exec(
		"INSERT INTO passwords (id, user_id, hashed_password) VALUES (?, ?, ?)",
		password.ID,
		password.UserID,
		password.HashedPassword,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pp passwordPersistence) FindPasswordByUserID(userID string) (*domain.Password, error) {
	password := domain.Password{}
	err := pp.QueryRow(
		"SELECT id, user_id, hashed_password FROM passwords WHERE user_id = ?",
		userID,
	).Scan(
		&password.ID,
		&password.UserID,
		&password.HashedPassword,
	)
	if err != nil {
		return &password, err
	}

	return &password, nil
}

func (pp passwordPersistence) UpdatePassword(userID, updatePassword string) error {
	_, err := pp.Exec(
		"UPDATE passwords SET hashed_password = ? WHERE user_id = ?",
		updatePassword,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
