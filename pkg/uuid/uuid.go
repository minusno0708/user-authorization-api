package uuid

import (
	"github.com/google/uuid"
)

func New() string {
	return uuid.NewString()
}

func IsUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
