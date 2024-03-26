package main

import (
	"trading-bot/model"
	"trading-bot/service"
)

func main() {
	svc := &service.BotService{}

	bot := &model.Bot{
		Svc: svc,
	}
	
	bot.Start()
}
