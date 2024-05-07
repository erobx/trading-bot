package main

import (
	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/types"
	"github.com/shopspring/decimal"
)

func main() {
	g := fillGroup()
	if g.IsReady() {
		g.TradeUp()
	}
}

func fillGroup() *model.Group {
	t := types.RandomTier()
	g := model.NewGroup(t)
	for i := 0; i < 10; i++ {
		s := buildSkin()
		g.AddSkin(s)
	}
	return g
}

func buildSkin() model.Skin {
	d, _ := decimal.NewFromString("10.12")
	price := types.DbDecimal(d)
	w := types.RandomWear()
	s := model.NewSkin("Redline", w, price)
	return s
}
