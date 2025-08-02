package user

import (
	"testing"

	"user-register-api/config"
	"user-register-api/domain/user"
	"user-register-api/pkg/errors"
	"user-register-api/pkg/uuid"
)

func Test_User_Insert_Success(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	tests := []struct {
		name     string
		userID   string
		username string
		email    string
	}{
		{
			name:     "正常系: 正常なユーザー",
			userID:   uuid.New(),
			username: uuid.New(),
			email:    uuid.New() + "@example.com",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			user, err := user.ReconstructUser(tt.userID, tt.username, tt.email)
			if err != nil {
				t.Fatalf("Failed to reconstruct user: %v", err)
			}

			err = repo.Insert(user)
			if !errors.Is(err, nil) {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}

func Test_User_Insert_Failed(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	insertUser, err := user.ReconstructUser(
		uuid.New(),
		uuid.New(),
		uuid.New()+"@example.com",
	)
	if err != nil {
		t.Fatalf("Failed to reconstruct user: %v", err)
	}

	err = repo.Insert(insertUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	tests := []struct {
		name        string
		userID      string
		username    string
		email       string
		expectError error
	}{
		{
			name:        "異常系: 重複するユーザーID",
			userID:      insertUser.ID(),
			username:    uuid.New(),
			email:       uuid.New() + "@example.com",
			expectError: errors.ErrConflict,
		},
		{
			name:        "異常系: 重複するユーザー名",
			userID:      uuid.New(),
			username:    insertUser.Username(),
			email:       uuid.New() + "@example.com",
			expectError: errors.ErrConflict,
		},
		{
			name:        "異常系: 重複するメールアドレス",
			userID:      uuid.New(),
			username:    uuid.New(),
			email:       insertUser.Email(),
			expectError: errors.ErrConflict,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			user, err := user.ReconstructUser(tt.userID, tt.username, tt.email)
			if err != nil {
				t.Fatalf("Failed to reconstruct user: %v", err)
			}

			err = repo.Insert(user)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}
		})
	}
}

func Test_User_FindByID(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	insertUser, err := user.ReconstructUser(
		uuid.New(),
		uuid.New(),
		uuid.New()+"@example.com",
	)
	if err != nil {
		t.Fatalf("Failed to reconstruct user: %v", err)
	}

	err = repo.Insert(insertUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	tests := []struct {
		name        string
		userID      string
		expectError error
	}{
		{
			name:        "正常系: 有効なユーザーID",
			userID:      insertUser.ID(),
			expectError: nil,
		},
		{
			name:        "異常系: 存在しないユーザーID",
			userID:      uuid.New(),
			expectError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			user, err := repo.FindByID(tt.userID)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}

			if tt.expectError == nil && user == nil {
				t.Error("Expected user to be found, but got nil")
			} else if tt.expectError == nil && user.ID() != tt.userID {
				t.Errorf("Expected user ID %s, got %s", tt.userID, user.ID())
			}
		})
	}
}

func Test_User_FindByUsername(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	insertUser, err := user.ReconstructUser(
		uuid.New(),
		uuid.New(),
		uuid.New()+"@example.com",
	)
	if err != nil {
		t.Fatalf("Failed to reconstruct user: %v", err)
	}

	err = repo.Insert(insertUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	tests := []struct {
		name        string
		username    string
		expectError error
	}{
		{
			name:        "正常系: 有効なユーザーID",
			username:    insertUser.Username(),
			expectError: nil,
		},
		{
			name:        "異常系: 存在しないユーザーID",
			username:    uuid.New(),
			expectError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			user, err := repo.FindByID(tt.username)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}

			if tt.expectError == nil && user == nil {
				t.Error("Expected user to be found, but got nil")
			} else if tt.expectError == nil && user.Username() != tt.username {
				t.Errorf("Expected user ID %s, got %s", tt.username, user.ID())
			}
		})
	}
}

func Test_User_Update(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	insertUser, err := user.ReconstructUser(
		uuid.New(),
		uuid.New(),
		uuid.New()+"@example.com",
	)
	if err != nil {
		t.Fatalf("Failed to reconstruct user: %v", err)
	}

	err = repo.Insert(insertUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	tests := []struct {
		name        string
		userID      string
		username    string
		email       string
		expectError error
	}{
		{
			name:        "正常系: usernameの更新",
			userID:      insertUser.ID(),
			username:    uuid.New(),
			email:       insertUser.Email(),
			expectError: nil,
		},
		{
			name:        "正常系: emailの更新",
			userID:      insertUser.ID(),
			username:    insertUser.Username(),
			email:       uuid.New() + "@example.com",
			expectError: nil,
		},
		{
			name:        "異常系: 存在しないユーザーID",
			userID:      uuid.New(),
			username:    uuid.New(),
			email:       uuid.New() + "@example.com",
			expectError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			user, err := user.ReconstructUser(tt.userID, tt.username, tt.email)
			if err != nil {
				t.Fatalf("Failed to reconstruct user: %v", err)
			}

			err = repo.Update(user)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}

			if tt.expectError == nil {
				updatedUser, err := repo.FindByID(tt.userID)
				if err != nil {
					t.Fatalf("Failed to find updated user: %v", err)
				}
				if updatedUser.Username() != tt.username {
					t.Errorf("Expected username %s, got %s", tt.username, updatedUser.Username())
				}
				if updatedUser.Email() != tt.email {
					t.Errorf("Expected email %s, got %s", tt.email, updatedUser.Email())
				}
			}
		})
	}
}

func Test_User_Delete(t *testing.T) {
	t.Parallel()

	db, err := config.ConnectDB("localhost:3306")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepository(db)

	insertUser, err := user.ReconstructUser(
		uuid.New(),
		uuid.New(),
		uuid.New()+"@example.com",
	)
	if err != nil {
		t.Fatalf("Failed to reconstruct user: %v", err)
	}

	err = repo.Insert(insertUser)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	tests := []struct {
		name        string
		userID      string
		expectError error
	}{
		{
			name:        "正常系: 有効なユーザーID",
			userID:      insertUser.ID(),
			expectError: nil,
		},
		{
			name:        "異常系: 存在しないユーザーID",
			userID:      uuid.New(),
			expectError: errors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err = repo.Delete(tt.userID)
			if !errors.Is(err, tt.expectError) {
				t.Errorf("Expected error %v, got %v", tt.expectError, err)
			}

			if tt.expectError == nil {
				deletedUser, err := repo.FindByID(tt.userID)
				if !errors.Is(err, errors.ErrNotFound) {
					t.Errorf("Expected user to be deleted, but found: %v", deletedUser)
				}
			}
		})
	}
}
