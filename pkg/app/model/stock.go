package model

import "github.com/erobx/trading-bot/pkg/types"

type Stock struct {
	Id     string
	Price  types.DbDecimal
	Amount int
}
