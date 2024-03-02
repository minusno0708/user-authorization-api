package repository

type TokenRepository interface {
	GenerateToken(userID, token string) error
	ValidateToken(token string) (string, error)
}
