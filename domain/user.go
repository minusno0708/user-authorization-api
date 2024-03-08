package domain

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(userID, username, password string) *User {
	if username == "" {
		username = userID
	}

	return &User{
		UserID:   userID,
		Username: username,
		Password: password,
	}
}
