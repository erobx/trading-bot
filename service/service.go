package service

import (
	"context"
)

type Service interface {
	GetSkinPrice(ctx context.Context) error
}

type BotService struct{}

func (bs *BotService) GetSkinPrice(ctx context.Context) error {
	return nil
}
