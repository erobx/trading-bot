package main

import (
	"errors"
	"sync"
)

type User struct {
	Balance float32
	Shares  []*Share
	svc     Service
	mu      sync.RWMutex
}

func NewUser(svc Service, b float32) *User {
	return &User{
		Balance: b,
		svc:     svc,
	}
}

func (u *User) FindSkin(name string) *Skin {
	return u.svc.FindSkin(name)
}

func (u *User) BuyShareOfSkin(name string) error {
	skin := u.svc.FindSkin(name)
	if skin == nil {
		return errors.New("could not find skin")
	}
	share := skin.buyShare()
	u.Shares = append(u.Shares, share)
	return nil
}

func (u *User) ListSkin(name string, price float32) error {
	return u.svc.ListSkin(name, price)
}

func (u *User) printShares() {
	for _, share := range u.Shares {
		share.printInfo()
	}
}
