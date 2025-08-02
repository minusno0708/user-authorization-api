package user

import (
	"testing"

	"user-register-api/config"
	passwordinfra "user-register-api/infra/password"
	userinfra "user-register-api/infra/user"
	"user-register-api/pkg/uuid"
)

func Test_User_Create(t *testing.T) {
	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Error(err)
	}

	userRepo := userinfra.NewUserRepository(db)
	passwordRepo := passwordinfra.NewPasswordRepository(db)
	userUseCase := NewUserUseCase(userRepo, passwordRepo)

	tests := []struct {
		name        string
		username    string
		email       string
		password    string
		expectError error
	}{
		{
			name:        "正常系: 有効なユーザー",
			username:    uuid.New(),
			email:       uuid.New() + "@example.com",
			password:    "P@ssw0rd",
			expectError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			user, err := userUseCase.Create(tt.username, tt.email, tt.password)
			if err != nil {
				if err.Error() != tt.expectError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectError, err)
				}
				return
			}

			if user.Username() != tt.username {
				t.Errorf("Expected username %s, got %s", tt.username, user.Username())
			}
			if user.Email() != tt.email {
				t.Errorf("Expected email %s, got %s", tt.email, user.Email())
			}
		})
	}
}

func Test_User_FindByUsername(t *testing.T) {
	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Error(err)
	}

	userRepo := userinfra.NewUserRepository(db)
	passwordRepo := passwordinfra.NewPasswordRepository(db)
	userUseCase := NewUserUseCase(userRepo, passwordRepo)

	createUser, err := userUseCase.Create(uuid.New(), uuid.New()+"@example.com", "P@ssw0rd")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		username    string
		expectError error
	}{
		{
			name:        "正常系: 有効なユーザー名",
			username:    createUser.Username(),
			expectError: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := userUseCase.FindByUsername(tt.username)
			if err != nil {
				if err.Error() != tt.expectError.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectError, err)
				}
				return
			}
		})
	}
}
