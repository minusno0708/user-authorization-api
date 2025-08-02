package me

import (
	"testing"
	"user-register-api/config"
	"user-register-api/domain/token"
	"user-register-api/domain/user"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"

	tokeninfra "user-register-api/infra/token"
	userinfra "user-register-api/infra/user"
)

func Test_Me_Get(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Error(err)
	}

	cdb, err := config.ConnectCacheDB("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}

	secretKey := "me_secret_key"
	userRepo := userinfra.NewUserRepository(db)
	tokenRepo := tokeninfra.NewTokenRepository(cdb)

	meUseCase := NewMeUseCase(secretKey, userRepo, tokenRepo)

	createUser, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatal(err)
	}

	err = userRepo.Insert(createUser)
	if err != nil {
		t.Fatal(err)
	}

	validToken := token.NewToken(createUser.ID())
	validTokenString, err := validToken.SignedString(secretKey)
	if err != nil {
		t.Fatal(err)
	}

	invalidToken := token.NewToken(createUser.ID())
	invalidTokenString, err := invalidToken.SignedString(secretKey)
	if err != nil {
		t.Fatal(err)
	}
	err = tokenRepo.Invalidate(invalidTokenString)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		tokenString string
		expectError error
	}{
		{
			name:        "正常系: 有効なトークンでユーザー情報を取得",
			tokenString: validTokenString,
			expectError: nil,
		},
		{
			name:        "異常系: 無効なトークンでユーザー情報を取得",
			tokenString: invalidTokenString,
			expectError: errors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			user, err := meUseCase.Get(tt.tokenString)
			if err != nil {
				if err.Error() != tt.expectError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectError, err)
				}
				return
			}

			if user == nil {
				t.Fatal("Expected user to be returned")
			}

			if user.ID() != createUser.ID() {
				t.Errorf("Expected user ID %s, got %s", createUser.ID(), user.ID())
			}
			if user.Username() != createUser.Username() {
				t.Errorf("Expected username %s, got %s", createUser.Username(), user.Username())
			}
			if user.Email() != createUser.Email() {
				t.Errorf("Expected email %s, got %s", createUser.Email(), user.Email())
			}
		})
	}
}

func Test_Me_Update(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Error(err)
	}

	cdb, err := config.ConnectCacheDB("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}

	secretKey := "me_secret_key"
	userRepo := userinfra.NewUserRepository(db)
	tokenRepo := tokeninfra.NewTokenRepository(cdb)

	meUseCase := NewMeUseCase(secretKey, userRepo, tokenRepo)

	createUser, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatal(err)
	}

	err = userRepo.Insert(createUser)
	if err != nil {
		t.Fatal(err)
	}

	validToken := token.NewToken(createUser.ID())
	validTokenString, err := validToken.SignedString(secretKey)
	if err != nil {
		t.Fatal(err)
	}

	invalidToken := token.NewToken(createUser.ID())
	invalidTokenString, err := invalidToken.SignedString(secretKey)
	if err != nil {
		t.Fatal(err)
	}
	err = tokenRepo.Invalidate(invalidTokenString)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name           string
		tokenString    string
		updateUsername string
		updateEmail    string
		expectError    error
	}{
		{
			name:           "正常系: 有効なトークンでユーザー情報を更新",
			tokenString:    validTokenString,
			updateUsername: uuid.New(),
			updateEmail:    uuid.New() + "@example.com",
			expectError:    nil,
		},
		{
			name:           "異常系: 無効なトークンでユーザー情報を更新",
			tokenString:    invalidTokenString,
			updateUsername: "",
			updateEmail:    "",
			expectError:    errors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			user, err := meUseCase.Update(tt.tokenString, tt.updateUsername, tt.updateEmail)
			if err != nil {
				if err.Error() != tt.expectError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectError, err)
				}
				return
			}

			if user == nil {
				t.Fatal("Expected user to be returned")
			}

			if user.ID() != createUser.ID() {
				t.Errorf("Expected user ID %s, got %s", createUser.ID(), user.ID())
			}
			if user.Username() != tt.updateUsername {
				t.Errorf("Expected username %s, got %s", tt.updateUsername, user.Username())
			}
			if user.Email() != tt.updateEmail {
				t.Errorf("Expected email %s, got %s", tt.updateEmail, user.Email())
			}
		})
	}
}

func Test_Me_Delete(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Error(err)
	}

	cdb, err := config.ConnectCacheDB("localhost:6379")
	if err != nil {
		t.Fatal(err)
	}

	secretKey := "me_secret_key"
	userRepo := userinfra.NewUserRepository(db)
	tokenRepo := tokeninfra.NewTokenRepository(cdb)

	meUseCase := NewMeUseCase(secretKey, userRepo, tokenRepo)

	createUser, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatal(err)
	}

	err = userRepo.Insert(createUser)
	if err != nil {
		t.Fatal(err)
	}

	validToken := token.NewToken(createUser.ID())
	validTokenString, err := validToken.SignedString(secretKey)
	if err != nil {
		t.Fatal(err)
	}

	invalidToken := token.NewToken(createUser.ID())
	invalidTokenString, err := invalidToken.SignedString(secretKey)
	if err != nil {
		t.Fatal(err)
	}
	err = tokenRepo.Invalidate(invalidTokenString)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		tokenString string
		expectError error
	}{
		{
			name:        "正常系: 有効なトークンでユーザー情報を更新",
			tokenString: validTokenString,
			expectError: nil,
		},
		{
			name:        "異常系: 無効なトークンでユーザー情報を更新",
			tokenString: invalidTokenString,
			expectError: errors.ErrForbidden,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := meUseCase.Delete(tt.tokenString)
			if err != nil {
				if err.Error() != tt.expectError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectError, err)
				}
				return
			}
		})
	}
}
