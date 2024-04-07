package main

import (
	"fmt"
	"math/rand"
	"trading-bot/model"
	"trading-bot/service"
)

func tradingBot() {
	marketChannel := make(chan string)
	go monitorMarket(marketChannel, 100)

	svc := service.NewBotService()
	svc = service.NewLogService(svc)

	bot := model.NewBot(svc, marketChannel)
	bot.Start()

}

func monitorMarket(c chan<- string, num int) {
	defer close(c)
	base := float32(5)

	for i := 0; i < num; i++ {
		increase := rand.Float32() < 0.7

		fluctuation := rand.Float32() * 1.5
		if !increase {
			fluctuation = -fluctuation
		}
		base += fluctuation

		if base < 0 {
			base = 0
		}
		moneyValue := fmt.Sprintf("$%.2f", base)
		c <- moneyValue
	}
}

func runSimulation() {
	marketService := NewMarketService()
	logger := NewLogService(marketService)
	user := NewUser(logger, 100)

	simulation := NewSimulation(user)
	simulation.start()

	// go simulation.GenerateValues()
	// go simulation.PrintFromChan()
	// select {}
}

type Simulation struct {
	simChan chan float32
	user    *User
}

func NewSimulation(user *User) *Simulation {
	return &Simulation{
		simChan: make(chan float32),
		user:    user,
	}
}

func (s *Simulation) start() {
	s.user.AddSkin("Redline", "Factory New", 30.21)
	s.user.AddSkin("Redline", "Factory New", 30.21)
	s.user.AddSkin("Redline", "Factory New", 10.99)
	s.user.GetSkin("Redline", "Factory New", 30.21)

	s.user.RemoveSkin("Redline", "Factory New", 30.21)
	s.user.RemoveSkin("Redline", "Factory New", 30.21)
	s.user.RemoveSkin("Redline", "Factory New", 30.21)
}

// func (s *Simulation) buyShares(name string, amount int) {
// 	for i := 0; i < amount; i++ {
// 		if err := s.user.BuyShareOfSkin(name); err != nil {
// 			fmt.Println(err)
// 		}
// 	}
// }

func (s *Simulation) GenerateValues() {
	for {
		s.simChan <- rand.Float32()
	}
}

func (s *Simulation) PrintFromChan() {
	for f := range s.simChan {
		fmt.Println(f)
	}
}
