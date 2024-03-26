package service

import (
	"context"
)

// Most popular AK-47, M4A1-S, AWP, USP-S and Glock descending

type Service interface {
	GetSkinPrice(ctx context.Context, id string) error
	WatchMarketValues(chan string) error
}
