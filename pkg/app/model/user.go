package model

type User struct {
	Balance float32
}

func NewUser(b float32) *User {
	return &User{
		Balance: b,
	}
}

// func (u *User) GetSkin(context context.Context, name, wear string) (Skin, error) {
// 	return u.svc.GetSkin(context, name, wear)
// }

// func (u *User) BuyShareOfSkin(context context.Context, name, wear string, price float32) error {
// 	_, err := u.svc.GetSkin(context, name, wear)
// 	// skin.buyShare()
// 	return err
// }

// func (u *User) AddSkin(context context.Context, name, wear string, price float32) error {
// 	return u.svc.AddSkin(context, name, wear, price)
// }

// func (u *User) RemoveSkin(context context.Context, name, wear string, price float32) error {
// 	return u.svc.RemoveSkin(context, name, wear, price)
// }
