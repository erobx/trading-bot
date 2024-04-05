package main

import (
	"fmt"
	"log"
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
	name := "Redline"
	s.buyShares(name, 1)
	s.user.printShares()

	err := s.user.ListSkin("Water Elemental", "Factory New", 32.45)
	// Should recover from error TODO
	if err != nil {
		log.Println(err)
	}

	skin := s.user.FindSkin("Water Elemental")
	fmt.Println("Name:", skin.Name)
}

func (s *Simulation) buyShares(name string, amount int) {
	for i := 0; i < amount; i++ {
		if err := s.user.BuyShareOfSkin(name); err != nil {
			fmt.Println(err)
		}
	}
}

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
