package repository

type TokenRepository interface {
	Invalidate(tokenString string) error
	IsValid(tokenString string) error
}
