package main

import (
	"errors"
)

type Service interface {
	FindSkin(name string) *Skin
	ListSkin(name string, price float32) error
}

type MarketService struct {
	market *Market
}

func NewMarketService() Service {
	return &MarketService{
		market: NewMarket(),
	}
}

func (ms *MarketService) FindSkin(name string) *Skin {
	skin := ms.market.Skins[name]
	return skin
}

func (ms *MarketService) ListSkin(name string, price float32) error {
	_, ok := ms.market.Skins[name]
	if ok {
		return errors.New("skin already exists")
	}
	skin := NewSkin(name, price)
	ms.market.Skins[name] = skin
	return nil
}

type Market struct {
	// direct db connection
	Skins map[string]*Skin
}

func NewMarket() *Market {
	return &Market{
		Skins: getSkins(),
	}
}

// Read from DB or file for list of skins
func getSkins() map[string]*Skin {
	skins := make(map[string]*Skin)
	skins["Redline"] = NewSkin("Redline", 10)
	return skins
}
