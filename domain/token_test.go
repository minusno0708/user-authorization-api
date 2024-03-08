package domain

import (
	"testing"
)

func TestDomainToken(t *testing.T) {
	userID := "test_user"
	token := NewToken(userID)
	tokenString, err := token.ToString()
	if err != nil {
		t.Error("Error while generating token")
	}
	if tokenString == "" {
		t.Error("TokenString is empty")
	}

	parsedToken, err := ParseToken(tokenString)
	if err != nil {
		t.Error("Error while parsing token")
	}
	if parsedToken.UserID() != userID {
		t.Error("UserID is not matched")
	}
	if parsedToken.UUID() == "" {
		t.Error("UUID is empty")
	}
}

func TestDomainTokenExpired(t *testing.T) {
	userID := "test_user"
	token := NewToken(userID)
	tokenString, err := token.ToString()
	if err != nil {
		t.Error("Error while generating token")
	}
	parsedToken, _ := ParseToken(tokenString)
	if parsedToken.IsExpired() {
		t.Error("Token is expired")
	}
}
