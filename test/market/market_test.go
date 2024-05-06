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

	err := m.AddSkin(s1)
	if err != nil {
		t.Error(err)
	}

	ok := m.UpdateStock(s1, true)
	if !ok {
		t.Errorf("somehow failing\n")
	}

	stock, ok := m.GetStock(names[0], wears[0])
	if !ok {
		t.Errorf("somehow failing\n")
	}

	if stock.Price.String() != "12.12" {
		t.Errorf("incorrect price, got %s", stock.Price.String())
	}
}

func TestAddSkins(t *testing.T) {
	p1, _ := decimal.NewFromString("69.42")
	p2, _ := decimal.NewFromString("42.42")
	p3, _ := decimal.NewFromString("50.00")

	s1 := model.NewSkin(names[1], wears[0], types.DbDecimal(p1))
	s2 := model.NewSkin(names[1], wears[0], types.DbDecimal(p2))
	s3 := model.NewSkin(names[1], wears[0], types.DbDecimal(p3))

	err := m.AddSkin(s1)
	if err != nil {
		t.Error(err)
	}
	
	err = m.AddSkin(s2)
	if err != nil {
		t.Error(err)
	}

	err = m.AddSkin(s3)
	if err != nil {
		t.Error(err)
	}
}

func TestGetSkin(t *testing.T) {
	skin, err := m.GetSkin(names[0], wears[0])
	if !err {
		t.Error(err)
	}

	if skin.Price.String() != "12.12" {
		t.Errorf("incorrect price, got %s", skin.Price.String())
	}
}

func TestGetSkins(t *testing.T) {
	skin, err := m.GetSkin(names[1], wears[0])
	if !err {
		t.Error(err)
	}

	if skin.Price.String() != "50" {
		t.Errorf("incorrect price, got %s\n", skin.Price.String())
	}
}

func TestRemoveSkin(t *testing.T) {
	err := m.RemoveSkin(names[0], wears[0])
	if !err {
		t.Error(err)
	}

	err = m.RemoveSkin(names[1], wears[0])
	if !err {
		t.Error(err)
	}

	skin, err := m.GetSkin(names[1], wears[0])
	if !err {
		t.Error(err)
	}

	if skin.Price.String() != "42.42" {
		t.Errorf("incorrect price, got %s", skin.Price.String())
	}
}

func TestGetStoc(t *testing.T) {
}