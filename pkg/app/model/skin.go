package model

import (
	"encoding/json"
	"fmt"

	"github.com/erobx/trading-bot/pkg/types"
)

type Skin struct {
	Name  string          `json:"Name"`
	Wear  string          `json:"Wear"`
	Price types.DbDecimal `json:"Price"`
}

func NewSkin(name, wear string, intial types.DbDecimal) Skin {
	return Skin{
		Name:  name,
		Wear:  wear,
		Price: intial,
	}
}

func (s Skin) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name  string
		Wear  string
		Price string
	}{
		Name:  s.Name,
		Wear:  s.Wear,
		Price: s.Price.String(),
	})
}

func (s *Skin) generateKey() string {
	return fmt.Sprintf("%s_%s", s.Name, s.Wear)
}
