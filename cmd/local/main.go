package main

import (
	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/types"
	"github.com/shopspring/decimal"
)

func main() {
	m, err := db.NewMarket()
	if err != nil {
		panic(err)
	}

	t := types.RandomTier()
	g := model.NewGroup(t)
	m.AddGroup(g)

	for i := 0; i < 10; i++ {
		s := buildSkin()
		m.AddSkin(s)
	}
}

func buildSkin() model.Skin {
	d, _ := decimal.NewFromString("10.12")
	price := types.DbDecimal(d)
	d, _ = decimal.NewFromString("0.00123")
	m := types.DbDecimal(d)
	d, _ = decimal.NewFromString("0.99")
	mx := types.DbDecimal(d)
	w := types.RandomWear()
	s := model.NewSkin("Redline", w, "AK-47", price, m, mx)
	return s
}
