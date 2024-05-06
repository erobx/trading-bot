package main

import (
	"testing"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/types"
	"github.com/shopspring/decimal"
)

var (
	m, _ = db.NewTestMarket()
	names = types.Names
	wears = types.Wears
)

func TestAddSkin(t *testing.T) {
	p1, _ := decimal.NewFromString("12.12")
	s1 := model.NewSkin(names[0], wears[0], types.DbDecimal(p1))

	addSkin(t, s1)
	getStock(t, names[0], wears[0], "12.12")
}

func TestAddSkins(t *testing.T) {
	p1, _ := decimal.NewFromString("69.42")
	p2, _ := decimal.NewFromString("42.42")
	p3, _ := decimal.NewFromString("50.00")

	s1 := model.NewSkin(names[1], wears[0], types.DbDecimal(p1))
	s2 := model.NewSkin(names[1], wears[0], types.DbDecimal(p2))
	s3 := model.NewSkin(names[1], wears[0], types.DbDecimal(p3))

	addSkin(t, s1)
	addSkin(t, s2)
	addSkin(t, s3)

	getSkin(t, names[1], wears[0], "50")
	getStock(t, names[1], wears[0], "50")
}

func TestGetSkin(t *testing.T) {
	getSkin(t, names[0], wears[0], "12.12")
}

func TestGetStock(t *testing.T) {
	getStock(t, names[0], wears[0], "12.12")
}

func TestGetSkins(t *testing.T) {
	getSkin(t, names[1], wears[0], "50")
}

func TestGetStocks(t *testing.T) {
	getStock(t, names[1], wears[0], "50")
}

func TestRemoveSkin(t *testing.T) {
	removeSkin(t, names[0], wears[0])
	removeSkin(t, names[1], wears[0])

	getSkin(t, names[1], wears[0], "42.42")
	getStock(t, names[1], wears[0], "42.42")
}

func TestAddAndRemove(t *testing.T) {
	p1, _ := decimal.NewFromString("10.11")
	p2, _ := decimal.NewFromString("11.11")
	p3, _ := decimal.NewFromString("12.11")
	p4, _ := decimal.NewFromString("13.11")
	p5, _ := decimal.NewFromString("14.11")

	s1 := model.NewSkin(names[2], wears[0], types.DbDecimal(p1))
	s2 := model.NewSkin(names[2], wears[0], types.DbDecimal(p2))
	s3 := model.NewSkin(names[2], wears[0], types.DbDecimal(p3))
	s4 := model.NewSkin(names[2], wears[0], types.DbDecimal(p4))
	s5 := model.NewSkin(names[2], wears[0], types.DbDecimal(p5))

	addSkin(t, s1)
	addSkin(t, s2)
	addSkin(t, s3)
	addSkin(t, s4)
	addSkin(t, s5)

	getStock(t, names[2], wears[0], "12.11")

	removeSkin(t, names[2], wears[0])
	removeSkin(t, names[2], wears[0])

	getStock(t, names[2], wears[0], "13.11")
}

func addSkin(t *testing.T, skin model.Skin) {
	err := m.AddSkin(skin)
	if err != nil {
		t.Error(err)
	}

	ok := m.UpdateStock(skin, true)
	if !ok {
		t.Errorf("somehow failing\n")
	}
}

func removeSkin(t *testing.T, name, wear string) {
	err := m.RemoveSkin(name, wear)
	if !err {
		t.Error(err)
	}

	skin := model.Skin{Name: name, Wear: wear}
	ok := m.UpdateStock(skin, false)
	if !ok {
		t.Errorf("couldn't update after removing\n")
	}
}

func getSkin(t *testing.T, name, wear, price string) {
	skin, err := m.GetSkin(name, wear)
	if !err {
		t.Error(err)
	}

	if skin.Price.String() != price {
		t.Errorf("incorrect price, got %s", skin.Price.String())
	}
}

func getStock(t *testing.T, name, wear, price string) {
	stock, ok := m.GetStock(name, wear)
	if !ok {
		t.Errorf("somehow failing\n")
	}

	if stock.Price.String() != price {
		t.Errorf("incorrect price, got %s", stock.Price.String())
	}
}

