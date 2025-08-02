package user

import (
	"database/sql"

	"user-register-api/domain/repository"
	"user-register-api/domain/user"
	"user-register-api/pkg/errors"
)

type userRepository struct {
	*sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db}
}

type userDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (ur userRepository) Insert(user *user.User) error {
	_, err := ur.Exec(
		"INSERT INTO users (id, username, email) VALUES (?, ?, ?)",
		user.ID(),
		user.Username(),
		user.Email(),
	)
	if err != nil {
		return errors.ErrConflict
	}

	return nil
}

func (ur userRepository) FindByID(userID string) (*user.User, error) {
	dbUser := userDTO{}
	err := ur.QueryRow(
		"SELECT id, username, email FROM users WHERE id = ? AND deleted_at IS NULL",
		userID,
	).Scan(
		&dbUser.ID,
		&dbUser.Username,
		&dbUser.Email,
	)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	domainUser, err := user.ReconstructUser(dbUser.ID, dbUser.Username, dbUser.Email)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return domainUser, nil
}

func (ur userRepository) FindByUsername(username string) (*user.User, error) {
	dbUser := userDTO{}
	err := ur.QueryRow(
		"SELECT id, username, email FROM users WHERE username = ? AND deleted_at IS NULL",
		username,
	).Scan(
		&dbUser.ID,
		&dbUser.Username,
		&dbUser.Email,
	)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	domainUser, err := user.ReconstructUser(dbUser.ID, dbUser.Username, dbUser.Email)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return domainUser, nil
}

func (ur userRepository) Update(user *user.User) error {
	result, err := ur.Exec(
		"UPDATE users SET username = ?, email = ? WHERE id = ?",
		user.Username(),
		user.Email(),
		user.ID(),
	)
	if err != nil {
		return errors.ErrInternalServer
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.ErrInternalServer
	}
	if rowsAffected == 0 {
		return errors.ErrNotFound
	}

	return nil
}

func (ur userRepository) Delete(userID string) error {
	result, err := ur.Exec(
		"DELETE FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.ErrInternalServer
	}
	if rowsAffected == 0 {
		return errors.ErrNotFound
	}

	return nil
}
