package main

type User struct {
	Balance    float32
	OwnedSkins map[string]map[string]*Skin
	svc        Service
}

func NewUser(svc Service, b float32) *User {
	return &User{
		Balance:    b,
		OwnedSkins: make(map[string]map[string]*Skin),
		svc:        svc,
	}
}

func (u *User) GetSkin(name, wear string, price float32) (Skin, error) {
	return u.svc.GetSkin(name, wear, price)
}

func (u *User) BuyShareOfSkin(name, wear string, price float32) error {
	_, err := u.svc.GetSkin(name, wear, price)
	// skin.buyShare()
	return err
}

func (u *User) AddSkin(name, wear string, price float32) error {
	return u.svc.AddSkin(name, wear, price)
}

func (u *User) RemoveSkin(name, wear string, price float32) error {
	return u.svc.RemoveSkin(name, wear, price)
}

// func (u *User) displayPositions() {
// 	for _, skin := range u.OwnedSkins {
// 		fmt.Printf("Name %s\n", skin.Name)
// 	}
// }
