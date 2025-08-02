package usecase

import (
	"testing"
	"user-register-api/config"
	"user-register-api/domain/password"
	"user-register-api/domain/user"
	passwordinfra "user-register-api/infra/password"
	tokeninfra "user-register-api/infra/token"
	userinfra "user-register-api/infra/user"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"
)

func Test_Auth_Login(t *testing.T) {
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
	passwordRepo := passwordinfra.NewPasswordRepository(db)
	tokenRepo := tokeninfra.NewTokenRepository(cdb)

	authUseCase := NewAuthUseCase(secretKey, userRepo, passwordRepo, tokenRepo)

	testuser, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if err := userRepo.Insert(testuser); err != nil {
		t.Fatal(err)
	}

	password, err := password.NewPassword(testuser.ID(), "P@ssw0rd")
	if err != nil {
		t.Fatal(err)
	}
	if err := passwordRepo.Insert(password); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name          string
		username      string
		rawPassword   string
		expectedError error
	}{
		{
			name:          "正常系: 正しいユーザー名とパスワードでログイン",
			username:      testuser.Username(),
			rawPassword:   "P@ssw0rd",
			expectedError: nil,
		},
		{
			name:          "異常系: 存在しないユーザー名でログイン",
			username:      "invaliduser",
			rawPassword:   "P@ssw0rd",
			expectedError: errors.ErrUnauthorized,
		},
		{
			name:          "異常系: 誤ったパスワードでログイン",
			username:      testuser.Username(),
			rawPassword:   "invalidpassword",
			expectedError: errors.ErrUnauthorized,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := authUseCase.Login(tt.username, tt.rawPassword)
			if err != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectedError, err)
				}
				return
			}
		})
	}
}

func Test_Auth_Logout(t *testing.T) {
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
	passwordRepo := passwordinfra.NewPasswordRepository(db)
	tokenRepo := tokeninfra.NewTokenRepository(cdb)

	authUseCase := NewAuthUseCase(secretKey, userRepo, passwordRepo, tokenRepo)

	testuser, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatal(err)
	}
	if err := userRepo.Insert(testuser); err != nil {
		t.Fatal(err)
	}

	password, err := password.NewPassword(testuser.ID(), "P@ssw0rd")
	if err != nil {
		t.Fatal(err)
	}
	if err := passwordRepo.Insert(password); err != nil {
		t.Fatal(err)
	}

	tokenString, err := authUseCase.Login(testuser.Username(), "P@ssw0rd")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name          string
		tokenString   string
		expectedError error
	}{
		{
			name:          "正常系: 正しいトークンでログアウト",
			tokenString:   tokenString,
			expectedError: nil,
		},
		{
			name:          "異常系: 無効なトークンでログアウト",
			tokenString:   "invalid_token",
			expectedError: errors.ErrUnauthorized,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := authUseCase.Logout(tt.tokenString)
			if err != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectedError, err)
				}
				return
			}
		})
	}
}
