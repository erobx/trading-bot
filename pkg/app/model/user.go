package model

type User struct {
	Balance   float32
	Inventory []Skin
}

func NewUser(b float32) *User {
	return &User{
		Balance: b,
	}
}

