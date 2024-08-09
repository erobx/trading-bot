package model

type Tradeup struct {
	Tier   string
	Active int
}

type DisplayTrade struct {
	TradeId int
	Tier    string
	Skins   []Skin
}

func NewTradeup(t string) Tradeup {
	return Tradeup{
		Tier:   t,
		Active: 1,
	}
}
