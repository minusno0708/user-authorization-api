package token

import (
	"testing"
	"time"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"

	"github.com/golang-jwt/jwt"
)

func Test_Token_NewToken(t *testing.T) {
	t.Parallel()

	userID := uuid.New()
	token := NewToken(userID)

	if token == nil {
		t.Error("Token is nil")
	}

	if token.UserID() != userID {
		t.Error("UserID is not matched")
	}

	if token.UUID() == "" {
		t.Error("UUID is empty")
	}

	if token.Exp() < time.Now().Unix() {
		t.Error("Token expiration time is over the current time")
	}
}

func Test_Token_Parse(t *testing.T) {
	t.Parallel()

	userID := uuid.New()
	token := NewToken(userID)
	signedToken, err := token.SignedString("secret")

	if err != nil {
		t.Error("Error while generating token")
	}

	tests := []struct {
		name          string
		validationKey string
		expectError   error
	}{
		{
			name:          "正常系: 有効なsecretキーで解析する",
			validationKey: "secret",
			expectError:   nil,
		},
		{
			name:          "異常系: 無効なsecretキーで解析する",
			validationKey: "invalid_key",
			expectError:   errors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			parsedToken, err := ParseToken(signedToken, tt.validationKey)

			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}

			if tt.expectError == nil {
				if parsedToken.UserID() != userID {
					t.Error("UserID is not matched")
				}
				if parsedToken.Exp() != token.Exp() {
					t.Error("Expiration time is not matched")
				}
			}
		})
	}
}

func Test_Token_Valid(t *testing.T) {
	t.Parallel()

	userID := uuid.New()

	tests := []struct {
		name        string
		token       *Token
		expectValid bool
	}{
		{
			name:        "正常系: 有効なトークンを検証する",
			token:       NewToken(userID),
			expectValid: true,
		},
		{
			name: "異常系: 期限切れのトークンを検証する",
			token: &Token{
				value: jwt.NewWithClaims(
					jwt.SigningMethodHS256,
					jwt.MapClaims{
						"user_id": userID,
						"uuid":    uuid.New(),
						"exp":     time.Now().Add(-time.Hour).Unix(),
					},
				),
			},
			expectValid: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.token.IsValid() != tt.expectValid {
				t.Errorf("Expected validity %v, got %v", tt.expectValid, tt.token.IsValid())
			}
		})
	}
}
