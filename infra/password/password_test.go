package password

import (
	"testing"

	"user-register-api/config"
	"user-register-api/domain/password"
	"user-register-api/domain/user"
	userinfra "user-register-api/infra/user"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"
)

func Test_Password_Insert_Success(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPasswordRepository(db)
	userRepo := userinfra.NewUserRepository(db)

	user, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	err = userRepo.Insert(user)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	tests := []struct {
		name     string
		userID   string
		password string
	}{
		{
			name:     "正常系: 正常なパスワード",
			userID:   user.ID(),
			password: "P@ssw0rd",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			password, err := password.NewPassword(tt.userID, tt.password)
			if err != nil {
				t.Fatalf("Failed to create password: %v", err)
			}
			err = repo.Insert(password)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func Test_Password_Insert_Failed(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPasswordRepository(db)
	userRepo := userinfra.NewUserRepository(db)

	user, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	err = userRepo.Insert(user)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	insertPassword, err := password.NewPassword(user.ID(), "P@ssw0rd")
	if err != nil {
		t.Fatalf("Failed to create password: %v", err)
	}
	err = repo.Insert(insertPassword)
	if err != nil {
		t.Fatalf("Failed to insert password: %v", err)
	}

	tests := []struct {
		name        string
		userID      string
		password    string
		expectError error
	}{
		{
			name:        "異常系: 重複するユーザーID",
			userID:      insertPassword.UserID(),
			password:    "P@ssw0rd",
			expectError: errors.ErrInternalServer,
		},
		{
			name:        "異常系: 存在しないユーザーID",
			userID:      uuid.New(),
			password:    "P@ssw0rd",
			expectError: errors.ErrInternalServer,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			password, err := password.NewPassword(tt.userID, tt.password)
			if err != nil {
				t.Fatalf("Failed to create password: %v", err)
			}

			err = repo.Insert(password)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_Password_FindByUserID(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPasswordRepository(db)
	userRepo := userinfra.NewUserRepository(db)

	user, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	err = userRepo.Insert(user)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	insertPassword, err := password.NewPassword(user.ID(), "P@ssw0rd")
	if err != nil {
		t.Fatalf("Failed to create password: %v", err)
	}
	err = repo.Insert(insertPassword)
	if err != nil {
		t.Fatalf("Failed to insert password: %v", err)
	}

	tests := []struct {
		name        string
		userID      string
		expectError error
	}{
		{
			name:        "正常系: 存在するユーザーID",
			userID:      user.ID(),
			expectError: nil,
		},
		{
			name:        "異常系: 存在しないユーザーID",
			userID:      uuid.New(),
			expectError: errors.ErrInternalServer,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			password, err := repo.FindByUserID(tt.userID)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}

			if tt.expectError == nil {
				if password.UserID() != tt.userID {
					t.Errorf("Expected userID %s, got %s", tt.userID, password.UserID())
				}
				if password.Compare("P@ssw0rd") != nil {
					t.Errorf("Expected password to match, but it did not")
				}
			}
		})
	}
}

func Test_Password_Update(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPasswordRepository(db)
	userRepo := userinfra.NewUserRepository(db)

	user, err := user.NewUser(uuid.New(), uuid.New()+"@example.com")
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
	err = userRepo.Insert(user)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	insertPassword, err := password.NewPassword(user.ID(), "P@ssw0rd")
	if err != nil {
		t.Fatalf("Failed to create password: %v", err)
	}
	err = repo.Insert(insertPassword)
	if err != nil {
		t.Fatalf("Failed to insert password: %v", err)
	}

	tests := []struct {
		name        string
		userID      string
		password    string
		expectError error
	}{
		{
			name:        "正常系: パスワードが更新できる",
			userID:      user.ID(),
			password:    "NewP@ssw0rd",
			expectError: nil,
		},
		{
			name:        "異常系: 存在しないユーザー",
			userID:      uuid.New(),
			password:    "NewP@ssw0rd",
			expectError: errors.ErrInternalServer,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			password, err := password.NewPassword(tt.userID, tt.password)
			if err != nil {
				t.Fatalf("Failed to create password: %v", err)
			}

			err = repo.Update(password)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}

			if tt.expectError == nil {
				updatedPassword, err := repo.FindByUserID(tt.userID)
				if err != nil {
					t.Fatalf("Failed to find updated password: %v", err)
				}
				if updatedPassword.UserID() != tt.userID {
					t.Errorf("Expected userID %s, got %s", tt.userID, updatedPassword.UserID())
				}
				if updatedPassword.Compare(tt.password) != nil {
					t.Errorf("Expected password to match, but it did not")
				}
			}
		})
	}
}
