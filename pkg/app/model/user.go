package model

type User struct {
	Balance float32
}

func NewUser(b float32) *User {
	return &User{
		Balance: b,
	}
}