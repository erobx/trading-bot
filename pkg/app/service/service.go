package service

import (
	"context"
	"fmt"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/types"
)

type Service interface {
	GetSkin(context context.Context, name, wear string) (model.Skin, error)
	AddSkin(context context.Context, name, wear string, price types.DbDecimal) error
	RemoveSkin(context context.Context, name, wear string, price types.DbDecimal) error
}

type MarketService struct {
	market *db.Market
}

func NewMarketService(m *db.Market) Service {
	return &MarketService{
		market: m,
	}
}

func (ms *MarketService) GetSkin(context context.Context, name, wear string) (model.Skin, error) {
	skin, ok := ms.market.GetSkin(name, wear)
	if !ok {
		return model.Skin{}, fmt.Errorf("skin not found")
	}
	return skin, nil
}

func (ms *MarketService) AddSkin(context context.Context, name, wear string, price types.DbDecimal) error {
	err := ms.market.AddSkin(model.NewSkin(name, wear, price))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ms *MarketService) RemoveSkin(context context.Context, name, wear string, price types.DbDecimal) error {
	if !ms.market.RemoveSkin(name, wear, price) {
		return fmt.Errorf("could not remove skin")
	}
	return nil
}
