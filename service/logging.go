package service

import (
	"context"
	"log"
)

type LogService struct {
	next Service
}

func NewLogService(next Service) Service {
	return &LogService{
		next: next,
	}
}

func (ls *LogService) GetSkinPrice(ctx context.Context, id string) error {
	log.Println("Logging skin price")
	return ls.next.GetSkinPrice(ctx, id)
}

func (ls *LogService) WatchMarketValues(mc chan string) error {
	log.Println("Watching market")
	return ls.next.WatchMarketValues(mc)
}