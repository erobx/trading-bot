package model

import (
	"github.com/erobx/trading-bot/pkg/types"
)

type Skin struct {
	Name       string          `json:"Name"`
	Weapon     string          `json:"Weapon"`
	Wear       string          `json:"Wear"`
	Color      string          `json:"Color"`
	Collection string          `json:"Collection"`
	FloatMin   types.DbDecimal `json:"Min"`
	FloatMax   types.DbDecimal `json:"Max"`
	Fl         types.DbDecimal `json:"Fl"`
}
