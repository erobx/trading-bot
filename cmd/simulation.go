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
