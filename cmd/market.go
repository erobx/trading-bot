package main

import (
	"fmt"
	"sync"
)

type Service interface {
	GetSkin(name, wear string, price float32) (Skin, error)
	AddSkin(name, wear string, price float32) error
	RemoveSkin(name, wear string, price float32) error
}

type MarketService struct {
	market *Market
}

func NewMarketService() Service {
	return &MarketService{
		market: NewMarket(),
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
	ms.market.AddSkin(NewSkin(name, wear, price))
	return nil
}

func (ms *MarketService) RemoveSkin(name, wear string, price float32) error {
	if !ms.market.RemoveSkin(name, wear, price) {
		return fmt.Errorf("could not remove skin")
	}
	return nil
}

type Market struct {
	mu    sync.RWMutex
	Skins map[string]map[Skin]int
}

func NewMarket() *Market {
	return &Market{
		Skins: make(map[string]map[Skin]int),
	}
}

func (m *Market) AddSkin(skin Skin) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := m.generateKey(skin)
	if m.Skins[key] == nil {
		m.Skins[key] = make(map[Skin]int)
	}
	m.Skins[key][skin]++
}

func (m *Market) GetSkin(name, wear string, price float32) (Skin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := m.generateKey(NewSkin(name, wear, price))
	if skinsMap, ok := m.Skins[key]; ok {
		for skin := range skinsMap {
			return skin, true
		}
	}
	return Skin{}, false
}

func (m *Market) RemoveSkin(name, wear string, price float32) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := m.generateKey(NewSkin(name, wear, price))
	if skinsMap, ok := m.Skins[key]; ok {
		for skin := range skinsMap {
			m.Skins[key][skin]--
			if m.Skins[key][skin] == 0 {
				delete(m.Skins[key], skin)
			}
			return true
		}
	}
	return false
}

func (m *Market) generateKey(skin Skin) string {
	return fmt.Sprintf("%s_%s_%.2f", skin.Name, skin.Wear, skin.Price)
}
