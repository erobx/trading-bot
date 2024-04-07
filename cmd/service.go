package main

import "fmt"

type Service interface {
	GetSkin(name, wear string, price float32) (Skin, error)
	AddSkin(name, wear string, price float32) error
	RemoveSkin(name, wear string, price float32) error
}

type MarketService struct {
	market *Market
}

func NewMarketService(m *Market) Service {
	return &MarketService{
		market: m,
	}
}

func (ms *MarketService) GetSkin(name, wear string, price float32) (Skin, error) {
	skin, ok := ms.market.GetSkin(name, wear, price)
	if !ok {
		return Skin{}, fmt.Errorf("skin not found")
	}
	return skin, nil
}

func (ms *MarketService) AddSkin(name, wear string, price float32) error {
	err := ms.market.AddSkin(NewSkin(name, wear, price))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ms *MarketService) RemoveSkin(name, wear string, price float32) error {
	if !ms.market.RemoveSkin(name, wear, price) {
		return fmt.Errorf("could not remove skin")
	}
	return nil
}
