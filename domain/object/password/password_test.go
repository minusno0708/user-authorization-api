package password

import (
	"testing"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"
)

func Test_Password_SetID(t *testing.T) {
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

			p := &Password{}
			err := p.SetID(tt.id)

			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_Password_SetUserID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		userID      string
		expectError error
	}{
		{
			name:        "正常系: 有効なユーザーIDを設定する",
			userID:      uuid.New(),
			expectError: nil,
		},
		{
			name:        "異常系: 空のユーザーIDを設定する",
			userID:      "",
			expectError: errors.ErrInternalServer,
		},
		{
			name:        "異常系: 無効なユーザーIDを設定する",
			userID:      "invalid-user-id",
			expectError: errors.ErrInternalServer,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Password{}
			err := p.SetUserID(tt.userID)

			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_Password_SetPassword(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		password    string
		expectError error
	}{
		{
			name:        "正常系: 有効なパスワードを設定する",
			password:    "P@ssw0rd",
			expectError: nil,
		},
		{
			name:        "異常系: 空のパスワードを設定する",
			password:    "",
			expectError: errors.ErrBadRequest,
		},
		{
			name:        "異常系: 7文字以下のパスワードを設定する",
			password:    "P@ssw0r",
			expectError: errors.ErrBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Password{}
			err := p.SetPassword(tt.password)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_Password_SetHashedPassword(t *testing.T) {
	t.Parallel()

	hashedPassword, err := hash("P@ssw0rd")
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	tests := []struct {
		name        string
		hashedPass  string
		expectError error
	}{
		{
			name:        "正常系: 有効なハッシュ化されたパスワードを設定する",
			hashedPass:  hashedPassword,
			expectError: nil,
		},
		{
			name:        "異常系: 空の文字列を設定する",
			hashedPass:  "",
			expectError: errors.ErrBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Password{}
			err := p.SetHashedPassword(tt.hashedPass)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_Password_Hash(t *testing.T) {
	t.Parallel()

	password := "P@ssw0rd"

	hashedPassword, err := hash(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	tests := []struct {
		name    string
		compare string
	}{
		{
			name:    "正常系: パスワードが変換前と一致しない",
			compare: password,
		},
		{
			name:    "正常系: パスワードが空文字ではない",
			compare: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if hashedPassword == tt.compare {
				t.Errorf("Expected hashed password is matched incorrect value, got %s", tt.compare)
			}
		})
	}
}

func Test_Password_Compare(t *testing.T) {
	t.Parallel()

	userID := uuid.New()
	correctPassword := "P@ssw0rd"

	p, err := NewPassword(userID, correctPassword)
	if err != nil {
		t.Fatalf("Failed to create password: %v", err)
	}

	tests := []struct {
		name            string
		comparePassword string
		expectError     error
	}{
		{
			name:            "正常系: パスワードが一致する",
			comparePassword: correctPassword,
			expectError:     nil,
		},
		{
			name:            "異常系: パスワードが一致しない",
			comparePassword: "wrong_password",
			expectError:     errors.ErrUnauthorized,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := p.Compare(tt.comparePassword)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}
