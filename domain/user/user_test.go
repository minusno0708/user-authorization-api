package user

import (
	"testing"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"
)

func Test_User_SetID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		id          string
		expectError error
	}{
		{
			name:        "正常系: 有効なUUIDを設定する",
			id:          uuid.New(),
			expectError: nil,
		},
		{
			name:        "異常系: 空のUUIDを設定する",
			id:          "",
			expectError: errors.ErrInternalServer,
		},
		{
			name:        "異常系: 無効なUUIDを設定する",
			id:          "invalid-uuid",
			expectError: errors.ErrInternalServer,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &User{}
			err := p.SetID(tt.id)

			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_User_SetUsername(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		username    string
		expectError error
	}{
		{
			name:        "正常系: 有効なユーザー名を設定する",
			username:    "username",
			expectError: nil,
		},
		{
			name:        "異常系: 空のユーザー名を設定する",
			username:    "",
			expectError: errors.ErrBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &User{}
			err := u.SetUsername(tt.username)

			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_User_SetEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		email       string
		expectError error
	}{
		{
			name:        "正常系: 有効なメールアドレスを設定する",
			email:       "test@example.com",
			expectError: nil,
		},
		{
			name:        "異常系: 空のメールアドレスを設定する",
			email:       "",
			expectError: errors.ErrBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &User{}
			err := u.SetEmail(tt.email)

			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}
