package repository

type TokenRepository interface {
	GenerateToken(userID, token string) error
	ValidateToken(userID string) (string, error)
}
