package password

import (
	"database/sql"

	"user-register-api/domain/password"
	"user-register-api/domain/repository"
	"user-register-api/pkg/errors"
)

type passwordRepository struct {
	*sql.DB
}

func NewPasswordRepository(db *sql.DB) repository.PasswordRepository {
	return &passwordRepository{db}
}

type passwordDTO struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	HashedPassword string `json:"hashed_password"`
}

func (pr passwordRepository) Insert(password *password.Password) error {
	_, err := pr.Exec(
		"INSERT INTO passwords (id, user_id, hashed_password) VALUES (?, ?, ?)",
		password.ID(),
		password.UserID(),
		password.HashedPassword(),
	)
	if err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

func (pr passwordRepository) FindByUserID(userID string) (*password.Password, error) {
	dbPassword := passwordDTO{}
	err := pr.QueryRow(
		"SELECT id, user_id, hashed_password FROM passwords WHERE user_id = ?",
		userID,
	).Scan(
		&dbPassword.ID,
		&dbPassword.UserID,
		&dbPassword.HashedPassword,
	)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	domainPassword, err := password.ReconstructPassword(
		dbPassword.ID,
		dbPassword.UserID,
		dbPassword.HashedPassword,
	)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	return domainPassword, nil
}

func (pr passwordRepository) Update(password *password.Password) error {
	result, err := pr.Exec(
		"UPDATE passwords SET hashed_password = ? WHERE user_id = ?",
		password.HashedPassword(),
		password.UserID(),
	)
	if err != nil {
		return errors.ErrInternalServer
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.ErrInternalServer
	}
	if rowsAffected != 1 {
		return errors.ErrInternalServer
	}

	return nil
}
