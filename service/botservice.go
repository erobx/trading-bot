package service

import (
	"context"
	"fmt"
)

type BotService struct {}

func NewBotService() Service {
	return &BotService{}
}

func (bs *BotService) GetSkinPrice(ctx context.Context, id string) error {

	return nil
}

func (bs *BotService) WatchMarketValues(mc chan string) error {
	for money := range mc {
		fmt.Printf("%s ", money)
	}
	return nil
}
