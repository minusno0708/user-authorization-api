package repository

type TokenRepository interface {
	SaveToken(userID, token string) error
	ValidateToken(userID string) (string, error)
	DeleteToken(userID string) error
}
