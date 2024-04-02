package main

import (
	"errors"
)

type User struct {
	Balance     float32
	Shares      []*Share
	marketStore Service
}

func NewUser(ms Service, b float32) *User {
	return &User{
		Balance:     b,
		marketStore: ms,
	}
}

func (u *User) BuyShareOfSkin(name string) error {
	skin := u.marketStore.FindSkin(name)
	if skin == nil {
		return errors.New("could not find skin")
	}
	share := skin.buyShare()
	u.Shares = append(u.Shares, share)
	return nil
}

func (u *User) printShares() {
	for _, share := range u.Shares {
		share.printInfo()
	}
}
