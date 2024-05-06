package sim

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/types"
)

var (
	names = types.Names
	wears = types.Wears
)

type Sim struct {
	market *db.Market
	done   chan struct{}
}

func NewSim(d chan struct{}) *Sim {
	m, err := db.NewMarket()
	if err != nil {
		panic(err)
	}

	return &Sim{
		market: m,
		done:   d,
	}
}

func (s *Sim) Start() {
	fmt.Printf("Sim is running...\n\n")
	skins := make(chan model.Skin)
	for {
		select {
		case <-s.done:
			fmt.Printf("Stopping sim...\n")
			return
		default:
			go s.chooseAction(skins)
			time.Sleep(time.Second * 1)
			skin := <-skins
			s.getStockPrice(skin.Name, skin.Wear)
		}
	}
}

func (s *Sim) getStockPrice(name, wear string) {
	fmt.Println("Logging stock price...")

	stock, ok := s.market.GetStock(name, wear)
	if !ok {
		fmt.Println("Error getting stock")
		return
	}
	fmt.Printf("Stock: %s $%s\n\n", stock.Id, stock.Price)
}

func (s *Sim) chooseAction(skins chan model.Skin) {
	r := rand.IntN(11)
	if r >= 4 {
		s.listSkin(skins)
	} else {
		s.delistSkin(skins)
	}
}

func (s *Sim) listSkin(skins chan model.Skin) {
	rs := randomSkin()
	fmt.Printf("Listing: %s, %s\n", rs.Name, rs.Wear)
	s.market.AddSkin(rs)
	s.updateStock(rs, true)
	skins <- rs
}

func (s *Sim) delistSkin(skins chan model.Skin) {
	rs := randomSkin()
	fmt.Printf("Delisting: %s, %s\n", rs.Name, rs.Wear)
	suc := s.market.RemoveSkin(rs.Name, rs.Wear)
	if !suc {
		fmt.Println("Error removing skin")
		return
	}
	s.updateStock(rs, false)
	skins <- rs
}

func (s *Sim) updateStock(skin model.Skin, check bool) {
	fmt.Println("Updating stock value...")
	suc := s.market.UpdateStock(skin, check)
	if !suc {
		fmt.Println("Couldn't update")
		return
	}
}

func randomSkin() model.Skin {
	prices := db.RandomPrices()

	nameIndex := rand.IntN(len(names))
	wearIndex := rand.IntN(len(wears))
	priceIndex := rand.IntN(len(prices))

	skin := model.Skin{
		Name:  names[nameIndex],
		Wear:  wears[wearIndex],
		Price: prices[priceIndex],
	}
	return skin
}
