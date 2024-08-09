package model

type User struct {
	Username  string
	Email     string
	Inventory []Skin
}

func NewUser(username, email, hash string, inv []Skin) *User {
	return &User{
		Username:  username,
		Email:     email,
		Inventory: inv,
	}
}
