package repository

type TokenRepository interface {
	SaveToken(userID, tokenUUID string) error
	ValidateToken(userID string) (string, error)
	DeleteToken(userID string) error
}
