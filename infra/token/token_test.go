package token

import (
	"testing"
	"user-register-api/config"
	"user-register-api/domain/token"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"
)

func Test_Token_Invalidate(t *testing.T) {
	t.Parallel()

	cdb, err := config.ConnectCacheDB("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewTokenRepository(cdb)

	tokenString, err := token.NewToken(uuid.New()).SignedString("secret")
	if err != nil {
		t.Fatal("Failed to create token:", err)
	}

	tests := []struct {
		name        string
		tokenString string
	}{
		{
			name:        "正常系: トークンを無効化する",
			tokenString: tokenString,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := repo.Invalidate(tt.tokenString)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func Test_Token_IsValid(t *testing.T) {
	t.Parallel()

	cdb, err := config.ConnectCacheDB("localhost:6379")
	if err != nil {
		t.Error(err)
	}

	repo := NewTokenRepository(cdb)

	invalidUserID := uuid.New()
	InvalidTokenString, err := token.NewToken(invalidUserID).SignedString("secret")
	if err != nil {
		t.Error(err)
	}
	err = repo.Invalidate(InvalidTokenString)
	if err != nil {
		t.Error("Failed to invalidate token:", err)
	}

	tests := []struct {
		name        string
		tokenString string
		expectError error
	}{
		{
			name:        "正常系: 無効化されたトークン",
			tokenString: InvalidTokenString,
			expectError: errors.ErrUnauthorized,
		},
		{
			name:        "正常系: 有効なトークン",
			tokenString: uuid.New(),
			expectError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err = repo.IsValid(tt.tokenString)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}
