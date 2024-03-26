package model

import (
	// "context"
	"trading-bot/service"
)

type Bot struct {
	svc           service.Service
	marketChannel chan string
}

func NewBot(s service.Service, mc chan string) *Bot {
	return &Bot{
		svc: s,
		marketChannel: mc,
	}
}

func (bot *Bot) Start() {
	bot.svc.WatchMarketValues(bot.marketChannel)
	// bot.svc.GetSkinPrice(context.TODO(), "1")
}
