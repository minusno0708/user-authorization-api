package repository

type TokenRepository interface {
	SaveToken(userID, tokenUuid string) error
	ValidateToken(userID string) (string, error)
	DeleteToken(userID string) error
}
