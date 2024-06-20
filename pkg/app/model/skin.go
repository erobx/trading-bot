package model

import (
	"encoding/json"
	"fmt"

	"github.com/erobx/trading-bot/pkg/types"
	"github.com/shopspring/decimal"
)

type Skin struct {
	Name  string          `json:"Name"`
	Wear  string          `json:"Wear"`
	Price types.DbDecimal `json:"Price"`
	Gun   string          `json:"Gun"`
	Min   types.DbDecimal `json:"Min"`
	Max   types.DbDecimal `json:"Max"`
}

func NewSkin(name, wear, gun string, price, m, mx types.DbDecimal) Skin {
	return Skin{
		Name:  name,
		Wear:  wear,
		Price: price,
		Min:   m,
		Max:   mx,
	}
}

func (s Skin) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name  string
		Wear  string
		Price string
		Gun   string
		Min   string
		Max   string
	}{
		Name:  s.Name,
		Wear:  s.Wear,
		Price: s.Price.String(),
		Gun:   s.Gun,
		Min:   s.Min.String(),
		Max:   s.Max.String(),
	})
}

func (s *Skin) generateKey() string {
	return fmt.Sprintf("%s_%s", s.Name, s.Wear)
}

func BuildSkin() Skin {
	d, _ := decimal.NewFromString("10.12")
	price := types.DbDecimal(d)
	d, _ = decimal.NewFromString("0.00123")
	m := types.DbDecimal(d)
	d, _ = decimal.NewFromString("0.99")
	mx := types.DbDecimal(d)
	w := types.RandomWear()
	s := NewSkin("Redline", w, "AK-47", price, m, mx)
	return s
}
