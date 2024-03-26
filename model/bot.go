package model

import (
	"context"
	"fmt"
	"trading-bot/service"
)

type Bot struct {
	Svc service.Service
}

func (bot *Bot) Start() {
	fmt.Println("Initializing")
	fmt.Println("Fetching skin prices")
	bot.Svc.GetSkinPrice(context.TODO())
}
