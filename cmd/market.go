package main

type Service interface {
	FindSkin(name string) *Skin
}

type MarketService struct {
	market *Market
}

func NewMarketService() Service {
	return &MarketService{
		market: NewMarket(),
	}
}

func (ms *MarketService) FindSkin(name string) *Skin {
	return ms.market.Skins[name]
}

type Market struct {
	// direct db connection
	Skins map[string]*Skin
}

func NewMarket() *Market {
	return &Market{
		Skins: getSkins(),
	}
}

// Read from DB or file for list of skins
func getSkins() map[string]*Skin {
	skins := make(map[string]*Skin)
	skins["Redline"] = NewSkin("Redline", 10)
	return skins
}

