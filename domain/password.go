package domain

type Password struct {
	value string
}

func NewPassword(pwd string) *Password {
	return &Password{value: pwd}
}
